import Search from './search.tsx';
import CommonHeader from './common.tsx';
import UserActions from './userActions.tsx';

function Header() {
  return (
    <CommonHeader>
      <Search />
      <div className='flex items-center justify-center gap-3 mr-2'>
        <UserActions />
      </div>
    </CommonHeader>
  );
}

export default Header;
