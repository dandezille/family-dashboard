require 'rails_helper'

RSpec.feature "ManageActivities", type: :feature do

  before(:each) do
    Rails.application.load_seed
  end

  scenario 'User creates an activity' do
    visit '/activities'
    click_on 'Add activity'

    attributes = attributes_for(:activity)
    fill_in 'activity[time]', with: attributes[:time]
    fill_in 'activity[symbol]', with: attributes[:symbol]
    fill_in 'activity[note]', with: attributes[:note]
    click_on 'Add'

    expect(page).to have_activity(attributes)
  end

  scenario 'User edits an activity' do
    activity = create(:activity)
    visit '/activities'

    within "[data-activity-id='#{activity.id}']" do
      click_on '✏️'
    end

    expect(page).to have_selector("input[value='#{activity.time.strftime("%H:%M")}']")
    expect(page).to have_selector("input[value='#{activity.symbol}']")
    expect(page).to have_selector("input[value='#{activity.note}']")

    attributes = attributes_for(:activity)
    fill_in 'activity[time]', with: attributes[:time]
    fill_in 'activity[symbol]', with: attributes[:symbol]
    fill_in 'activity[note]', with: attributes[:note]
    click_on 'Change'

    expect(page).to have_activity(attributes)
  end

  scenario 'User deletes an activity' do
    activity = create(:activity)
    visit '/activities'

    within "[data-activity-id='#{activity.id}']" do
      click_on '🗑️'
    end

    expect(page).not_to have_text(activity.time.strftime("%H:%M"))
    expect(page).not_to have_text(activity.symbol)
    expect(page).not_to have_text(activity.note)
  end

  def have_activity(params)
    have_text(params[:time].strftime("%H:%M")) &&
      have_text(params[:symbol]) &&
      have_text(params[:note])
  end

  def edit_link_for(activity)
    find("[data-activity-id='#{activity.id}']").find('a', text: '✏️')
  end
end
