import { PropsWithChildren } from 'react';
import { useNavigate } from 'react-router-dom';
import { Image, Text } from '@fluentui/react-components';

import { appConfig } from '../../constants.ts';

function CommonHeader({ children }: PropsWithChildren) {
  const navigate = useNavigate();

  return (
    <div className='w-full h-[40px] px-2 flex items-center justify-between border-b-[1px] border-[#eaeaea] mb-[40px] absolute'>
      <div
        onClick={() => navigate('/app')}
        className='h-full mb-0 flex items-center cursor-pointer gap-2'
      >
        <Image src='/favicon.ico' height={30} width={30} className='' />

        <div className='flex flex-col'>
          <Text
            as='strong'
            style={{
              padding: 0,
              marginTop: -2,
              marginBottom: -2,
              fontSize: '0.9rem',
            }}
          >
            {appConfig.name}
          </Text>

          <Text
            style={{
              padding: 0,
              marginTop: -5,
              marginBottom: -2,
              fontSize: '0.7rem',
            }}
          >
            {appConfig.tagline}
          </Text>
        </div>
      </div>
      {children}
    </div>
  );
}
export default CommonHeader;
