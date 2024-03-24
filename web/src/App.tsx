import { Fragment, Suspense, lazy } from 'react';
import { Toaster } from '@fluentui/react-components';
import { QueryParamProvider } from 'use-query-params';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import { ReactQueryDevtools } from '@tanstack/react-query-devtools';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { ReactRouter6Adapter } from 'use-query-params/adapters/react-router-6';

import Loader from './components/loader.tsx';
import UiContextProvider from './contexts/ui.tsx';
import { TOASTER_ID } from './hooks/notifications.tsx';

const Home = lazy(() => import('./pages/home.tsx'));

const ReactQueryDevtoolsProduction = lazy(() =>
  import('@tanstack/react-query-devtools/production').then((d) => ({
    default: d.ReactQueryDevtools,
  }))
);

const queryClient = new QueryClient();

function App() {
  return (
    <Suspense fallback={<Loader />}>
      <QueryClientProvider client={queryClient}>
        {!import.meta.env.PROD ? (
          <Fragment>
            <ReactQueryDevtools initialIsOpen={false} />
            <ReactQueryDevtoolsProduction />
          </Fragment>
        ) : null}

        <UiContextProvider>
          <Toaster id={TOASTER_ID} />

          <div className='w-screen h-screen'>
            <BrowserRouter>
              <QueryParamProvider adapter={ReactRouter6Adapter}>
                <Routes>
                  <Route path='' Component={Home} />
                </Routes>
              </QueryParamProvider>
            </BrowserRouter>
          </div>
        </UiContextProvider>
      </QueryClientProvider>
    </Suspense>
  );
}

export default App;
