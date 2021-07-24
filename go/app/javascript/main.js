import ActivitiesController from './activities_controller.js'
import LinkController from './link_controller.js'
import TimeController from './time_controller.js'
import WeatherController from './weather_controller.js'

(() => {
  const application = Stimulus.Application.start()
  application.register("activities", ActivitiesController)
  application.register("link", LinkController)
  application.register("time", TimeController)
  application.register("weather", WeatherController)
})()
