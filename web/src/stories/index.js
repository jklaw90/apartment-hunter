import React from 'react';
import 'bootstrap/dist/css/bootstrap.css';

import { storiesOf } from '@storybook/react';

import App from '../App.js';
import Example from '../components/Button.js';

storiesOf('Major', module)
  .add('app', () => <App />)
;

storiesOf('Button', module)
  .add('button', () => <Example/>)
;
