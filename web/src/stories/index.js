import React from 'react';

import {storiesOf} from '@storybook/react';

import App from '../App.js';
import ListingRow from '../components/ListingRow.js';

storiesOf('Major', module).add('App', () => <App />);

storiesOf('Listings', module).add('ListingRow', () => <ListingRow />);
