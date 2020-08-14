import React, { useState, useEffect } from 'react';
import moment from 'moment';

import Activities from './Activities';
import ActivityDisplay from './Activity';
import NextActivity from './NextActivity';
import TimeDisplay from './Time';
import Weather, { GetWeather } from './Weather';

interface Props {
  activities: Activities;
  get_weather: GetWeather;
}

function useTime() {
  const [time, set_time] = useState(moment());

  useEffect(() => {
    const id = setInterval(() => { set_time(moment()); }, 1000);
    return () => clearInterval(id);
  });

  return time;
}

export default function App(props: Props) {
  const time = useTime();

  const activity = props.activities.current(time);
  const next_activity = props.activities.next(time);

  return (
    <div
      style={{
        display: 'flex',
        flex: 1,
        minHeight: '100vh',
      }}
    >
      <ActivityDisplay activity={activity} time={time} />
      <div
        style={{
          display: 'flex',
          flexDirection: 'column',
          justifyContent: 'space-between',
          alignItems: 'flex-end',
          padding: '15px',
        }}
      >
        <div>
          <TimeDisplay time={time} />
          <Weather
            get_weather={props.get_weather}
            update_interval={5 * 60 * 1000}
          />
        </div>
        <NextActivity activity={next_activity.symbol} />
      </div>
    </div>
  );
}
