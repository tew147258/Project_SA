import { createPlugin } from '@backstage/core';
import ConfirmationUI from './components/ConfirmationUI';
import StadiumUI from './components/StadiumUI';
import Login from './components/Login';

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', Login);
    router.registerRoute('/confirmationui', ConfirmationUI);
    router.registerRoute('/stadiumui', StadiumUI);
  },
});
