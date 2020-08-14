import React from 'react';
import ReactDOM from 'react-dom';

import App from './App';

import { create_activities } from './Activities';
import { get_open_weather_map_data } from './Weather';

document.addEventListener('DOMContentLoaded', () => {
  ReactDOM.render(
    <React.StrictMode>
      <App
        get_activities={() => Promise.resolve(create_activities())}
        get_weather={get_open_weather_map_data}
      />
    </React.StrictMode>,
    document.getElementById('dashboard')
  );
});
