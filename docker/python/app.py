#!/usr/bin/python
import os
from BaseHTTPServer import BaseHTTPRequestHandler,HTTPServer

SERVER_PORT = int(os.environ['SERVER_PORT'])

#This class will handles any incoming request from
#the browser 
class myHandler(BaseHTTPRequestHandler):
	
	#Handler for the GET requests
	def do_GET(self):
		self.send_response(200)
		self.send_header('Content-type','text/html')
		self.end_headers()
		# Send the html message
		self.wfile.write("Hello World")
		return

try:
	#Create a web server and define the handler to manage the
	#incoming request
	server = HTTPServer(('', SERVER_PORT), myHandler)
	print 'Started httpserver on port ' , SERVER_PORT
	
	#Wait forever for incoming htto requests
	server.serve_forever()

except KeyboardInterrupt:
	print '^C received, shutting down the web server'
	server.socket.close()
