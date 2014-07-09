#!/usr/bin/env ruby

require 'open-uri'

if Time.now.friday?
  URI.parse("http://hazlo.herokuapp.com/emails/weekly").open
end
