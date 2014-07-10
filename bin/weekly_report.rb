#!/usr/bin/env ruby

require 'open-uri'

if Time.now.monday?
  URI.parse("http://hazlo.herokuapp.com/emails/weekly").open
end
