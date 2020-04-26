import moment from "moment";

interface Activity {
  start: moment.Moment;
  duration: number;
  symbol: string;
}

interface Activities {
  current(time: moment.Moment): Activity;
}

class ActivitiesImplementation implements Activities {
  constructor(activities: { [time: string]: string }) {
    this.activities = activities;
  }

  private current_activity_index(time_strings: string[], time: moment.Moment) {
    const times = time_strings.map((t) => moment(t, "HH:mm"));

    for (var i = 0; i < times.length - 1; i++) {
      if (times[i] <= time && times[i + 1] > time) {
        return i;
      }
    }

    return times.length - 1;
  }

  private time_diff(current: moment.Moment, next: moment.Moment) {
    if (next > current) {
      return next.diff(current);
    }

    return next.add(1, "day").diff(current);
  }

  current(time: moment.Moment) {
    const times = Object.keys(this.activities);

    const current_index = this.current_activity_index(times, time);
    const current_start = this.parseTime(times[current_index]);

    const next_index = (current_index + 1) % times.length;
    const next_start = this.parseTime(times[next_index]);

    const activity_duration = this.time_diff(current_start, next_start);

    return {
      start: current_start,
      duration: activity_duration,
      symbol: this.activities[times[current_index]],
    };
  }

  private activities: { [time: string]: string };

  private parseTime(time: string) {
    return moment(time, "HH:mm");
  }

  private nextTimeIndex(times: string[], current_time: moment.Moment) {
    const index = times.findIndex((k) => this.parseTime(k) > current_time);
    if (index !== -1) {
      return index;
    }

    return times.length; // Handle last activity
  }
}

const default_activities = {
  "07:30": "🥐", // breakfast
  "08:15": "🦷", // Teeth and face wash
  "08:20": "📚", // French story and drawing
  "08:30": "🎨", // Craft (colouring, painting, drawing)
  "09:00": "🍳", // Baking
  "09:30": "️🚶‍♂️", // Walk
  "10:00": "️☕", // Coffee and listen
  "10:30": "️🧩", // Play
  "12:00": "️📺", // Cartoon
  "12:30": "️🍽️", // Lunch and listen
  "13:00": "️️️📺", // Cartoon
  "13:30": "️️️🧩️", // Play
  "13:50": "️️️📖️", // Story
  "14:00": "️️️️🛏️", // Nap
  "16:30": "️️️🧩️", // Play
  "17:00": "️🚶‍♂️", // Walk
  "17:30": "️️️📺", // Cartoon
  "18:00": "️🍽️", // Dinner and listen
  "18:30": "️️️📺", // Cartoon
  "19:00": "️️️🛀", // Bath or shower
  "19:30": "🦷", // Teeth and face wash
  "19:40": "️️️📖️", // Story
  "19:45": "🛏️", // Bed
};

export function create_activities(
  activities: { [time: string]: string } = default_activities
) {
  return new ActivitiesImplementation(activities);
}

export default Activities;
