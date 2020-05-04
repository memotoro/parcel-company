#!/usr/bin/env ruby

require 'json'
require 'rest-client'

command = ARGV.first

class Trucks
  class << self
    def base_path
      'http://localhost:8080'
    end

    def create
      counter = 1
      while counter <= 10
        data = { model: "Model #{counter}", capacityKg: counter * 10 }
        result = RestClient.post("#{base_path}/trucks", data.to_json, { content_type: :json })
        counter += 1
      end
    end

    def load
      counter = 1
      while counter <= 10
        truck_id = rand(1..10)
        data = { weightKg: counter * 0.5 }
        result = RestClient.post("#{base_path}/trucks/#{truck_id}/parcels", data.to_json, { content_type: :json })
        counter += 1
      end
    end
  end
end

case command
when 'create'
  Trucks.create
when 'load'
  Trucks.load
when 'unload'
  Trucks.unload
when 'list'
  Trucks.list
end
