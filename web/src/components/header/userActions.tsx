import {
  Avatar,
  Divider,
  Menu,
  MenuItem,
  MenuList,
  MenuPopover,
  MenuTrigger,
  Persona,
} from '@fluentui/react-components';
import { truncate } from 'lodash-es';
import { SignOut20Regular } from '@fluentui/react-icons';

import useAuth from '../../hooks/auth.tsx';

type Item = {
  key: string;
  icon: any;
  label: string;
  onClick: () => void;
};

const items: Array<Item> = [];

function UserActions() {
  const { user, logout } = useAuth();

  return (
    <Menu>
      <MenuTrigger disableButtonEnhancement>
        <Avatar />
      </MenuTrigger>

      <MenuPopover>
        <MenuList style={{ minWidth: 200, padding: 8 }}>
          <Persona
            name={user?.name}
            secondaryText={truncate(user?.email, { length: 20 })}
          />

          <Divider className='my-2' />

          <MenuItem onClick={logout} icon={<SignOut20Regular />}>
            Logout
          </MenuItem>

          {items.map((item) => (
            <MenuItem
              key={item.key}
              icon={<item.icon />}
              onClick={item.onClick}
            >
              {item.label}
            </MenuItem>
          ))}
        </MenuList>
      </MenuPopover>
    </Menu>
  );
}

export default UserActions;
