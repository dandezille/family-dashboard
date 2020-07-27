import React from "react";
import ReactDOM from "react-dom";

import { create_test_weather_provider, WeatherData } from "../models/WeatherProvider";

import Weather from "./Weather";

it("renders successfully", async () => {
  const data: WeatherData = {
    temp: 20,
  };

  const div = document.createElement("div");
  ReactDOM.render(
    <Weather
      weather_provider={create_test_weather_provider(data)}
      update_interval={5 * 1000}
    />,
    div
  );
});