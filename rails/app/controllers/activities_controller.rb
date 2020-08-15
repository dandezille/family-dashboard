class ActivitiesController < ApplicationController
  def index
    render json: {
      '08:00': '⏰', # wake up
      '08:15': '🥐', # breakfast
      '08:45': '🦷', # Teeth and face wash
      '09:00': '🇫🇷', # French story and drawing
      '09:15': '️🚶', # Walk
      '10:00': '️☕',  # Coffee and listen
      '10:30': '️🧩', # Play
      '11:45': '️📺', # Cartoon
      '12:15': '️🍽️', # Lunch and listen
      '13:15': '️️️📖️', # Story
      '13:30': '️️️️🛏️', # Np
      '15:00': '️️️📖️', # Story
      '15:15': '️️️🧩️', # Play
      '16:30': '️🚶', # Walk
      '17:00': '️️️🧩️', # Play
      '18:00': '️️️📺', # Cartoon
      '18:30': '️🍽️', # Dinner and listen
      '19:15': '️️️🛀', # Bath or shower
      '19:45': '🦷', # Teeth and face wash
      '20:00': '🛏️', # Bed
    }
  end
end
