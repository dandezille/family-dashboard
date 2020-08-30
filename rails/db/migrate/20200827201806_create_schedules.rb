class CreateSchedules < ActiveRecord::Migration[6.0]
  def change
    create_table :schedules do |t|
      t.string :name

      t.timestamps
    end
    add_index :schedules, :name, unique: true
  end
end
