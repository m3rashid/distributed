import { Suspense, lazy } from 'react';
import { Toaster } from '@fluentui/react-components';
import { QueryParamProvider } from 'use-query-params';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import { ReactRouter6Adapter } from 'use-query-params/adapters/react-router-6';

import Loader from './components/loader';
import UiContextProvider from './contexts/ui';
import { TOASTER_ID } from './hooks/notifications';

const Home = lazy(() => import('./pages/home'));

function App() {
  return (
    <Suspense fallback={<Loader />}>
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
    </Suspense>
  );
}

export default App;
