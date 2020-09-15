from __future__ import absolute_import
from __future__ import division
from __future__ import unicode_literals

from rasa_sdk import Action
from rasa_sdk.events import SlotSet, Restarted, AllSlotsReset
import zomatopy
import json

import smtplib
from email.mime.text import MIMEText
from email.mime.multipart import MIMEMultipart

class ActionSearchRestaurants(Action):
	def name(self):
		return 'action_search_restaurants'
		
	def run(self, dispatcher, tracker, domain):
		cnt = 0
		config={ "user_key":"e65d2d421f76e6521d4bb9758237b1ef"}
		zomato = zomatopy.initialize_app(config)
		loc = tracker.get_slot('location')

		cuisine = tracker.get_slot('cuisine')
		price = tracker.get_slot('price')
#		print(price)
		location_detail=zomato.get_location(loc, 1)
		d1 = json.loads(location_detail)
		lat=d1["location_suggestions"][0]["latitude"]
		lon=d1["location_suggestions"][0]["longitude"]
		cuisines_dict={'american':1,'chinese':25,'italian':55,'north indian':50,'south indian':85}
		price_dict={'low':1, 'mid':2, 'high':3}
		results=zomato.restaurant_search("", lat, lon, str(cuisines_dict.get(cuisine)), 100)
		d = json.loads(results)
		response=""
		response_list=[]
		if d['results_found'] == 0:
			response= "no results"
		else:
			for restaurant in sorted(d['restaurants'], key=lambda x: x['restaurant']['user_rating']['aggregate_rating'], reverse=True):
				#response=response+ "Found "+ restaurant['restaurant']['name']+ " in "+ restaurant['restaurant']['location']['address']+"\n"
				if((restaurant['restaurant']['average_cost_for_two'] <= 300) and (price_dict.get(price) == 1) and (cnt < 10)):
					amount = str(restaurant['restaurant']['average_cost_for_two'])
					response_list.append(restaurant['restaurant']['name']+ " "+ restaurant['restaurant']['location']['address']+ " has been rated "+ restaurant['restaurant']['user_rating']['aggregate_rating']+ " with average rate for 2 people as "+ str(restaurant['restaurant']['average_cost_for_two']))
					cnt += 1
				elif((restaurant['restaurant']['average_cost_for_two'] > 300) and (restaurant['restaurant']['average_cost_for_two'] <= 700) and (price_dict.get(price) == 2) and (cnt < 10)):
					amount = str(restaurant['restaurant']['average_cost_for_two'])
					response_list.append(restaurant['restaurant']['name']+ " "+ restaurant['restaurant']['location']['address']+ " has been rated "+ restaurant['restaurant']['user_rating']['aggregate_rating']+ " with average rate for 2 people as "+ str(restaurant['restaurant']['average_cost_for_two']))
					cnt += 1
				elif((restaurant['restaurant']['average_cost_for_two'] > 700) and (price_dict.get(price) == 3) and (cnt < 10)):
					amount = str(restaurant['restaurant']['average_cost_for_two'])
					response_list.append(restaurant['restaurant']['name']+ " "+ restaurant['restaurant']['location']['address']+ " has been rated "+ restaurant['restaurant']['user_rating']['aggregate_rating']+ " with average rate for 2 people as "+ str(restaurant['restaurant']['average_cost_for_two']))
					cnt += 1

				response_str = "\n".join(response_list[:5])
				response = str(response_str)

			response_mail = "\n".join(response_list)
			response_mail = str(response_mail)
											
			if(cnt == 0):
				response = "No restaurants found matching your criteria. Do you want us to look for other restaurants?"
				dispatcher.utter_message(response)
			else:
				#print("hello_dispatch")
				dispatcher.utter_message(response)
				
				#return [SlotSet('location',loc)]
			return [SlotSet('response_list', response_mail)]

class ActionLocationValidation(Action):
	def name(self):
		return 'action_chk_location'

	def run(self, dispatcher, tracker, domain):
		loc = tracker.get_slot('location')
		loc = loc.lower()

		avail_loc = ['Ahmedabad', 'Bengaluru', 'Bangalore', 'Chennai', 'Delhi', 'Hyderabad', 'Kolkata' , 'Mumbai', 'Thane', 'Muppini', 'Agra', 'Ajmer', 
           		     'Aligarh', 'Amravati', 'Amritsar', 'Asansol', 'Aurangabad', 'Bareilly', 'Belgaum', 'Bhavnagar', 'Bhiwandi', 'Bhopal', 
                     'Bhubaneswar', 'Bikaner', 'Bilaspur', 'Bokaro Steel City', 'Chandigarh', 'Coimbatore', 'Cuttack', 'Dehradun', 'Dhanbad', 
                     'Bhilai', 'Durgapur', 'Dindigul', 'Erode', 'Faridabad', 'Firozabad', 'Ghaziabad', 'Gorakhpur', 'Gulbarga', 'Guntur', 'Gwalior', 
                     'Gurgaon', 'Guwahati', 'Hamirpur', 'Hubli–Dharwad', 'Indore', 'Jabalpur', 'Jaipur', 'Jalandhar', 'Jammu', 'Jamnagar', 
                     'Jamshedpur', 'Jhansi', 'Jodhpur', 'Kakinada', 'Kannur', 'Kanpur', 'Kochi', 'Kolhapur', 'Kollam', 'Kozhikode', 'Kurnool', 
                     'Ludhiana', 'Lucknow', 'Madurai', 'Malappuram', 'Mathura', 'Mangalore', 'Meerut', 'Moradabad', 'Mysore', 'Nagpur', 'Nanded', 
                     'Nashik', 'Nellore', 'Noida', 'Patna', 'Pondicherry', 'Purulia', 'Prayagraj', 'Raipur', 'Rajkot', 'Rajahmundry', 'Ranchi', 
                     'Rourkela', 'Salem', 'Sangli', 'Shimla', 'Siliguri', 'Solapur', 'Srinagar', 'Surat', 'Thanjavur', 'Thiruvananthapuram', 
                     'Thrissur', 'Tiruchirappalli', 'Tirunelveli', 'Ujjain', 'Bijapur', 'Vadodara', 'Varanasi', 'Vasai-Virar City', 
                     'Vijayawada', 'Visakhapatnam', 'Vellore', 'Warangal']

		out = map(lambda x:x.lower(), avail_loc) 
		avail_loc = list(out)

		if loc not in avail_loc:
			dispatcher.utter_message('Sorry!!! We do not operate in this area. Do you want us to look for other cities?')
			location_check = 'invalid'
		else:
			location_check = 'valid'

		return [SlotSet('location_validity', location_check)]

class ActionSendEmail(Action):
	def name(self):
		return 'action_send_email'

	def run(self, dispatcher, tracker, domain):
		from_user = 'anand.venkatupgrad@gmail.com'
		to_user = tracker.get_slot('email')
		password = 'Upgrad@2020'
		s = smtplib.SMTP('smtp.gmail.com',587)
		s.starttls()
		s.login(from_user, password)
		subject = 'Foodie App'
		msg = MIMEMultipart()
		msg['From'] = from_user
		msg['TO'] = to_user
		msg['Subject'] = subject
		body = tracker.get_slot('response_list')
		msg.attach(MIMEText(body,'plain'))
		text = msg.as_string()
		s.sendmail(from_user,to_user,text)
		s.close()

		return [AllSlotsReset()]

class ActionRestart(Action):
	def name(self):
		return 'action_restart'

	def run(self, dispatcher, tracker, domain):
		return[Restarted()]

class ActionSlotReset(Action):
	def name(self):
		return 'action_slot_reset'

	def run(self, dispatcher, tracker, domain):
		return[AllSlotsReset()]




		