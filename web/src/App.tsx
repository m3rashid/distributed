import { Fragment, Suspense, lazy } from 'react';
import { QueryParamProvider } from 'use-query-params';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import { ReactQueryDevtools } from '@tanstack/react-query-devtools';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { ReactRouter6Adapter } from 'use-query-params/adapters/react-router-6';

import UiContextProvider from './contexts/ui.tsx';

const Home = lazy(() => import('./pages/home.tsx'));

const ReactQueryDevtoolsProduction = lazy(() =>
  import('@tanstack/react-query-devtools/production').then((mod) => ({
    default: mod.ReactQueryDevtools,
  }))
);

const queryClient = new QueryClient();

function App() {
  return (
    <UiContextProvider>
      <Suspense
        fallback={
          <div className='flex items-center justify-center h-screen'>
            {/* <Spinner size='extra-large' /> */}
          </div>
        }
      >
        <QueryClientProvider client={queryClient}>
          {!import.meta.env.PROD ? (
            <Fragment>
              <ReactQueryDevtools initialIsOpen={false} />
              <ReactQueryDevtoolsProduction />
            </Fragment>
          ) : null}

          {/* <Toaster id={TOASTER_ID} /> */}

          <div className='w-screen h-screen'>
            <BrowserRouter>
              <QueryParamProvider adapter={ReactRouter6Adapter}>
                <Routes>
                  <Route path='' Component={Home} />
                </Routes>
              </QueryParamProvider>
            </BrowserRouter>
          </div>
        </QueryClientProvider>
      </Suspense>
    </UiContextProvider>
  );
}

export default App;
