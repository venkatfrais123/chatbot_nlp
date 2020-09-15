## complete path
* greet
    - utter_greet
* restaurant_search
    - utter_ask_location
* restaurant_search{"location": "delhi"}
    - slot{"location": "delhi"}
    - action_chk_location
* restaurant_search{"location_validity": "valid"}
    - slot{"location_validity": "valid"}
    - utter_ask_cuisine
* restaurant_search{"cuisine": "chinese"}
    - slot{"cuisine": "chinese"}
    - utter_ask_budget
* restaurant_search{"price": "low"}
    - slot{"price": "low"}
    - action_search_restaurants
    - slot{"response_list": "Momo Sa-Khang by Kailash Kitchen 2/247, Perinbavilas Complex, Opposite Chitra Apartments, Choolaimedu, Chennai has been rated 4.6\nMezze Abhinav Center, Old 19, New 4, Co-operative Colony, Off Chamiers Road, Alwarpet, Chennai has been rated 4.6\nFusilli Reasons 2/64, Ormes Main Road, Kilpauk, Chennai has been rated 4.5\nYaa Mohaideen Biryani Old Shop 4/158, Church Road, Opposite Uzhavar Santhai, Pallavaram, Chennai has been rated 4.4\nWriter's Cafe 98, Peter's Road, Near Philip's Service Centre, Gopalapuram, Chennai has been rated 4.3\nYaa Mohaideen Briyani 336 & 338, Main Road, Pallavaram, Chennai has been rated 4.3"}
    - utter_email_check
* get_email_ind{"email_ind": "Send Email"}
    - slot{"email_ind": "Send Email"}
    - utter_ask_email
* get_email{"email": "anand19.r@gmail.com"}
    - slot{"email": "anand19.r@gmail.com"}
    - action_send_email
    - utter_email_confirm
* affirm
    - utter_goodbye
    - action_restart
## interactive_story_1
* greet
    - utter_greet
* restaurant_search{"cuisine": "south indian", "location": "chennai"}
    - slot{"cuisine": "south indian"}
    - slot{"location": "chennai"}
    - action_chk_location
    - slot{"location_validity": "valid"}
    - utter_ask_budget
* restaurant_search{"price": "mid"}
    - slot{"price": "mid"}
    - action_search_restaurants
    - slot{"response_list": "Prems Graama Bhojanam 23/10, 3rd Main Road, Kasturba Nagar, Adyar, Chennai has been rated 4.6\nThat Mallu Joint 2nd Floor, PMG Complex, GN Chetty Road, T. Nagar, Chennai has been rated 4.4\nNellai Vairamaligai 60, Thirumalai Pillai Road, Rama Kamath Puram, T. Nagar, Chennai has been rated 4.2\nMana Andhra 18/24, Khader Nawaskhan Road, Nungambakkam, Chennai has been rated 4.2\nKovai Alankar Vilas AC 3, 2nd Avenue, Anna Nagar East, Chennai . has been rated 4.1\nShree Konar Vilas 128, Perambur Barracks Road, Purasavakkam, Chennai has been rated 4.1"}
    - utter_email_check
* get_email_ind{"email_ind": "Send Email"}
    - slot{"email_ind": "Send Email"}
    - utter_ask_email
* get_email{"email": "venkat.frais@gmail.com"}
    - slot{"email": "venkat.frais@gmail.com"}
    - action_send_email
    - utter_email_confirm
* goodbye
    - utter_goodbye
    - action_restart

## interactive_story_1
* greet
    - utter_greet
* restaurant_search
    - utter_ask_location
* restaurant_search{"location": "Madurai"}
    - slot{"location": "Madurai"}
    - action_chk_location
    - slot{"location_validity": "valid"}
    - utter_ask_cuisine
* restaurant_search{"cuisine": "Italian"}
    - slot{"cuisine": "Italian"}
    - utter_ask_budget
* restaurant_search{"price": "low"}
    - slot{"price": "low"}
    - action_search_restaurants
    - slot{"response_list": "Kumar Mess 57, Alagar Kovil Road, Tallakulam, KK Nagar, Madurai has been rated 4.7\nZaitoon 32, 80 Feet Road, Anna Nagar, Sathamangalam, Madurai has been rated 4.6\nKFC 61/4B, Pattu Arcade, Bypass Road, Opposite World Of Titan, SS Colony, Madurai has been rated 4.4\nKFC 64, 80 Feet Road, Anna Nagar, Sathamangalam, Madurai has been rated 4.4\nDomino's Pizza Near Aristo Hospital, Bypass Main Road, SS Colony, Madurai has been rated 4.4\nDomino's Pizza Ground Floor, Sivasakthi Plaza, 80 Feet Road, Managairi, KK Nagar, Madurai has been rated 4.3\nAl Ahad Muslim Biriyani 5 /12, Subburaman Street, Gandhi Nagar, Anna Bus Stand, Anna Nagar, Sathamangalam, Madurai has been rated 4.3\nMukku Kadai K Subbu Moongil Thottam Bharathi Street, KK Nagar, Madurai has been rated 4.2\nExpress Briyani 135, 2nd East Main Road, Anna Nagar, Sathamangalam, Madurai has been rated 4.1\nDindigul Thalappakatti 1, Melur Main Road, Opposite to District Court, KK Nagar, Madurai has been rated 4.0"}
    - utter_email_check
* get_email_ind{"email_ind": "Don't send Email"}
    - slot{"email_ind": "Don't send Email"}
    - utter_goodbye
    - action_restart

## interactive_story_1
* greet
    - utter_greet
* restaurant_search{"cuisine": "cafe", "location": "chennai"}
    - slot{"cuisine": "cafe"}
    - slot{"location": "chennai"}
    - action_chk_location
    - slot{"location_validity": "valid"}
    - utter_ask_budget
* restaurant_search{"price": "high"}
    - slot{"price": "high"}
    - action_search_restaurants
    - slot{"response_list": "Zuka 100/1, Dr Radhakrishnan Salai, 8th Street, Radha Krishnan Corner, Mylapore, Chennai has been rated 4.3\nThis Way In 9/5, Dheenadayalu Street, T. Nagar, Chennai has been rated 4.2\n80 Degrees East 34/27, 28th Street, Thillai Ganga Nagar, Nanganallur, Chennai has been rated 4.0\nStarbucks Coffee TBR Tower, Y Block, 2nd Avenue, Near Ayyappa Temple, Anna Nagar East, Chennai has been rated 4.0\nThe Coffee Bean & Tea Leaf 1st Floor, Phoenix Market City, Velachery Main Road, Velachery, Chennai has been rated 4.0\nStarbucks Coffee Ground Floor, Bergamo Mall, Khader Nawaz Khan Road, Nungambakkam, Chennai has been rated 4.0\nStarbucks Coffee Ground Floor, Express Avenue Mall, Whites Road, Royapettah, Chennai has been rated 4.0\nUrban Desi House 1/343, 1st Main Road, Sri Sai Nagar, 200 Feet Road,  Thuraipakkam, Chennai has been rated 3.9\nStarbucks Coffee 92/5, Ground Floor, MGN Splendor, Off Chamiers Road, Alwarpet, Chennai has been rated 3.9\nStarbucks Coffee Ground Floor, PMG Complex, GN Chetty Road, T. Nagar, Chennai has been rated 3.9"}
    - utter_email_check
* get_email_ind{"email_ind": "Don't send Email"}
    - slot{"email_ind": "Don't send Email"}
    - utter_goodbye
    - action_restart

## interactive_story_1
* greet
    - utter_greet
* restaurant_search{"price": "high", "cuisine": "north indian", "location": "mumbai"}
    - slot{"cuisine": "north indian"}
    - slot{"location": "mumbai"}
    - slot{"price": "high"}
    - action_chk_location
    - slot{"location_validity": "valid"}
    - action_search_restaurants
    - slot{"response_list": "Butter Chicken Factory Shop 7, Silver Cross Building, 16th & 33rd Road Junction, Pali Hill, Bandra West, Mumbai has been rated 4.5\nThe Sahib Room & Kipling Bar - The St. Regis Mumbai The St. Regis, 462, Level 9M, Senapati Bapat Marg, Lower Parel, Mumbai has been rated 4.5\nTanatan Ramee Guestline Hotel, 462, A.B. Nair Road, Juhu, Mumbai has been rated 4.4\nVeranda 331, Dr Ambedkar Road, Pali Hill, Bandra West, Mumbai has been rated 4.3\nSpiceKlub 8A, Janta Industrial Estate, Senapati Bapat Road, Opposite Phoenix Mills, Lower Parel, Mumbai has been rated 4.3\nZaffran Todi & Company, Mathuradas Mills Compound, Lower Parel, Mumbai has been rated 4.2\nKoyla Above Gulf Hotel, HNA Azmi Marg, Colaba, Mumbai has been rated 4.1\nBombay Brasserie Villa 12-A, Lotus Cinema Building, Opposite Nehru Centre, Lotus Colony, Worli, Mumbai has been rated 4.1\nTaftoon Bar & Kitchen Unit 2, C-31 & 32, G Block, Naman Centre, Opposite Dena Bank, Bandra Kurla Complex, Mumbai has been rated 4.1\nCopper Chimney Door 12A, Ground Floor, Lotus Court, Dr. Annie Besant Road, Worli, Mumbai has been rated 4.1"}
    - utter_email_check
* get_email_ind{"email_ind": "Send Email"}
    - slot{"email_ind": "Send Email"}
    - utter_ask_email
* get_email{"email": "venkat.frais@gmail.com"}
    - slot{"email": "venkat.frais@gmail.com"}
    - action_send_email
    - utter_email_confirm
* affirm
    - utter_goodbye
    - action_restart

## interactive_story_1
* greet
    - utter_greet
* restaurant_search{"price": "high", "location": "mumbai"}
    - slot{"location": "mumbai"}
    - slot{"price": "high"}
    - action_chk_location
    - slot{"location_validity": "valid"}
    - utter_ask_cuisine
* restaurant_search{"cuisine": "Italian"}
    - slot{"cuisine": "Italian"}
    - action_search_restaurants
    - slot{"response_list": "Butter Chicken Factory Shop 7, Silver Cross Building, 16th & 33rd Road Junction, Pali Hill, Bandra West, Mumbai has been rated 4.5\nThe Sahib Room & Kipling Bar - The St. Regis Mumbai The St. Regis, 462, Level 9M, Senapati Bapat Marg, Lower Parel, Mumbai has been rated 4.5\nTanatan Ramee Guestline Hotel, 462, A.B. Nair Road, Juhu, Mumbai has been rated 4.4\nVeranda 331, Dr Ambedkar Road, Pali Hill, Bandra West, Mumbai has been rated 4.3\nSpiceKlub 8A, Janta Industrial Estate, Senapati Bapat Road, Opposite Phoenix Mills, Lower Parel, Mumbai has been rated 4.3\nZaffran Todi & Company, Mathuradas Mills Compound, Lower Parel, Mumbai has been rated 4.2\nKoyla Above Gulf Hotel, HNA Azmi Marg, Colaba, Mumbai has been rated 4.1\nBombay Brasserie Villa 12-A, Lotus Cinema Building, Opposite Nehru Centre, Lotus Colony, Worli, Mumbai has been rated 4.1\nTaftoon Bar & Kitchen Unit 2, C-31 & 32, G Block, Naman Centre, Opposite Dena Bank, Bandra Kurla Complex, Mumbai has been rated 4.1\nCopper Chimney Door 12A, Ground Floor, Lotus Court, Dr. Annie Besant Road, Worli, Mumbai has been rated 4.1"}
    - utter_email_check
* get_email_ind{"email_ind": "Send Email"}
    - slot{"email_ind": "Send Email"}
    - utter_ask_email
* get_email{"email": "venkat.frais@gmail.com"}
    - slot{"email": "venkat.frais@gmail.com"}
    - action_send_email
    - utter_email_confirm
* affirm
    - utter_goodbye
    - action_restart

## interactive_story_1
* greet
    - utter_greet
* restaurant_search{"price": "low", "location": "mumbai"}
    - slot{"location": "mumbai"}
    - slot{"price": "high"}
    - action_chk_location
    - slot{"location_validity": "valid"}
    - utter_ask_cuisine
* restaurant_search{"cuisine": "American"}
    - slot{"cuisine": "American"}
    - action_search_restaurants
    - slot{"response_list": "Butter Chicken Factory Shop 7, Silver Cross Building, 16th & 33rd Road Junction, Pali Hill, Bandra West, Mumbai has been rated 4.5\nThe Sahib Room & Kipling Bar - The St. Regis Mumbai The St. Regis, 462, Level 9M, Senapati Bapat Marg, Lower Parel, Mumbai has been rated 4.5\nTanatan Ramee Guestline Hotel, 462, A.B. Nair Road, Juhu, Mumbai has been rated 4.4\nVeranda 331, Dr Ambedkar Road, Pali Hill, Bandra West, Mumbai has been rated 4.3\nSpiceKlub 8A, Janta Industrial Estate, Senapati Bapat Road, Opposite Phoenix Mills, Lower Parel, Mumbai has been rated 4.3\nZaffran Todi & Company, Mathuradas Mills Compound, Lower Parel, Mumbai has been rated 4.2\nKoyla Above Gulf Hotel, HNA Azmi Marg, Colaba, Mumbai has been rated 4.1\nBombay Brasserie Villa 12-A, Lotus Cinema Building, Opposite Nehru Centre, Lotus Colony, Worli, Mumbai has been rated 4.1\nTaftoon Bar & Kitchen Unit 2, C-31 & 32, G Block, Naman Centre, Opposite Dena Bank, Bandra Kurla Complex, Mumbai has been rated 4.1\nCopper Chimney Door 12A, Ground Floor, Lotus Court, Dr. Annie Besant Road, Worli, Mumbai has been rated 4.1"}
    - utter_email_check
* get_email_ind{"email_ind": "Send Email"}
    - slot{"email_ind": "Send Email"}
    - utter_ask_email
* get_email{"email": "venkat.frais@gmail.com"}
    - slot{"email": "venkat.frais@gmail.com"}
    - action_send_email
    - utter_email_confirm
* affirm
    - utter_goodbye
    - action_restart

## interactive_story_1
* greet
    - utter_greet
* restaurant_search{"location": "Faridabad"}
    - slot{"location": "Faridabad"}
    - action_chk_location
    - slot{"location_validity": "valid"}
    - utter_ask_cuisine
* restaurant_search{"cuisine": "North Indian"}
    - slot{"cuisine": "North Indian"}
    - utter_ask_budget
* restaurant_search{"price": "high"}
    - slot{"price": "high"}
    - action_search_restaurants
    - slot{"response_list": "Saravana Bhavan P-13, Connaught Circus, Connaught Place, New Delhi has been rated 4.6\nAMA Cafe House 6, New Colony, Majnu ka Tilla, New Delhi has been rated 4.5\nTriveni Terrace Cafe 205, Tansen Marg, Opposite FICCI Auditorium, Mandi House, New Delhi has been rated 4.5\nPunjabi Angithi 32-22, A 4, DDA Market, Paschim Vihar, New Delhi has been rated 4.3"}
    - utter_email_check
* get_email_ind{"email_ind": "Send Email"}
    - slot{"email_ind": "Send Email"}
    - utter_ask_email
* get_email{"email": "venkat.frais@gmail.com"}
    - slot{"email": "venkat.frais@gmail.com"}
    - action_send_email
    - utter_email_confirm
* goodbye
    - utter_goodbye
    - action_restart

## interactive_story_1
* greet
    - utter_greet
* restaurant_search{"price": "high", "cuisine": "north indian", "location": "mumbai"}
    - slot{"cuisine": "north indian"}
    - slot{"location": "mumbai"}
    - slot{"price": "high"}
    - action_chk_location
    - slot{"location_validity": "valid"}
    - action_search_restaurants
    - slot{"response_list": "Butter Chicken Factory Shop 7, Silver Cross Building, 16th & 33rd Road Junction, Pali Hill, Bandra West, Mumbai has been rated 4.5\nThe Sahib Room & Kipling Bar - The St. Regis Mumbai The St. Regis, 462, Level 9M, Senapati Bapat Marg, Lower Parel, Mumbai has been rated 4.5\nTanatan Ramee Guestline Hotel, 462, A.B. Nair Road, Juhu, Mumbai has been rated 4.4\nVeranda 331, Dr Ambedkar Road, Pali Hill, Bandra West, Mumbai has been rated 4.3\nSpiceKlub 8A, Janta Industrial Estate, Senapati Bapat Road, Opposite Phoenix Mills, Lower Parel, Mumbai has been rated 4.3\nZaffran Todi & Company, Mathuradas Mills Compound, Lower Parel, Mumbai has been rated 4.2\nKoyla Above Gulf Hotel, HNA Azmi Marg, Colaba, Mumbai has been rated 4.1\nBombay Brasserie Villa 12-A, Lotus Cinema Building, Opposite Nehru Centre, Lotus Colony, Worli, Mumbai has been rated 4.1\nTaftoon Bar & Kitchen Unit 2, C-31 & 32, G Block, Naman Centre, Opposite Dena Bank, Bandra Kurla Complex, Mumbai has been rated 4.1\nCopper Chimney Door 12A, Ground Floor, Lotus Court, Dr. Annie Besant Road, Worli, Mumbai has been rated 4.1"}
    - utter_email_check
* get_email_ind{"email_ind": "Send Email"}
    - slot{"email_ind": "Send Email"}
    - utter_ask_email
* get_email{"email": "venkat.frais@gmail.com"}
    - slot{"email": "venkat.frais@gmail.com"}
    - action_send_email
    - utter_email_confirm
* affirm
    - utter_goodbye
    - action_restart

## interactive_story_1
* greet
    - utter_greet
* restaurant_search{"location": "Faridabad"}
    - slot{"location": "Faridabad"}
    - action_chk_location
    - slot{"location_validity": "valid"}
    - utter_ask_cuisine
* restaurant_search{"cuisine": "North Indian"}
    - slot{"cuisine": "North Indian"}
    - utter_ask_budget
* restaurant_search{"price": "mid"}
    - slot{"price": "mid"}
    - action_search_restaurants
    - slot{"response_list": "Saravana Bhavan P-13, Connaught Circus, Connaught Place, New Delhi has been rated 4.6\nAMA Cafe House 6, New Colony, Majnu ka Tilla, New Delhi has been rated 4.5\nTriveni Terrace Cafe 205, Tansen Marg, Opposite FICCI Auditorium, Mandi House, New Delhi has been rated 4.5\nPunjabi Angithi 32-22, A 4, DDA Market, Paschim Vihar, New Delhi has been rated 4.3"}
    - utter_email_check
* get_email_ind{"email_ind": "Send Email"}
    - slot{"email_ind": "Send Email"}
    - utter_ask_email
* get_email{"email": "venkat.frais@gmail.com"}
    - slot{"email": "venkat.frais@gmail.com"}
    - action_send_email
    - utter_email_confirm
* goodbye
    - utter_goodbye
    - action_restart

## interactive_story_1
* greet
    - utter_greet
* restaurant_search{"cuisine": "mexican", "location": "mumbai"}
    - slot{"cuisine": "mexican"}
    - slot{"location": "mumbai"}
    - action_chk_location
    - slot{"location_validity": "valid"}
    - utter_ask_budget
* restaurant_search{"price": "high"}
    - slot{"price": "high"}
    - action_search_restaurants
    - slot{"response_list": "Bastian B/1, New Kamal Building, Opposite National College, Linking Road, Bandra West, Mumbai has been rated 4.8\nAer - Four Seasons Four Seasons Hotel, 1/136, E Moses Road, Worli, Mumbai has been rated 4.7\nSammy Sosa Shop 18, Meera CHS, Near Mega Mall, Oshiwara Link Road, Oshiwara, Andheri West has been rated 4.6\nThe Nutcracker Modern House, Dr. V.B. Gandhi Marg, Kala Ghoda, Opposite One Forbes Building, Fort, Mumbai has been rated 4.6\nColaba Social 24, Ground Floor, Glen Rose Building, Behind Taj Mahal Palace, Apollo Bunder, Colaba, Mumbai has been rated 4.4\n1 BHK - Brew House Kitchen Ground Floor, Near Meera Towers CHS, Oshiwara Link Road, Oshiwara, Andheri West, Mumbai has been rated 4.4\nGymkhana 91 Bar & Kitchen 1st Floor, Raghuvanshi Mills Compound, Senapati Bapat Marg, Lower Parel, Mumbai has been rated 4.4\nThe Bombay Canteen Plot C 68, Unit 2, Jet Airways - Godrej, Bandra Kurla Complex, Mumbai has been rated 4.3\n145 Kala Ghoda 145, Kala Ghoda, Fort, Mumbai has been rated 4.3\nJamjar Diner 7A & B, Near Washing Bay, Yari Road, Arram Nagar 2, Versova, Andheri West, Mumbai has been rated 4.3"}
    - utter_email_check
* get_email_ind{"email_ind": "Send Email"}
    - slot{"email_ind": "Send Email"}
    - utter_ask_email
* get_email{"email": "venkat.frais@gmail.com"}
    - slot{"email": "venkat.frais@gmail.com"}
    - action_send_email
    - utter_email_confirm
* affirm
    - utter_goodbye
    - action_restart

## interactive_story_1
* greet{"greet": "Heya"}
    - utter_greet
* restaurant_search{"location": "New York"}
    - slot{"location": "New York"}
    - action_chk_location
    - slot{"location_validity": "invalid"}
* goodbye
    - utter_goodbye
    - action_restart

## interactive_story_1
* greet{"greet": "Heya"}
    - utter_greet
* restaurant_search{"location": "New York"}
    - slot{"location": "New York"}
    - action_chk_location
    - slot{"location_validity": "invalid"}
* goodbye
    - utter_goodbye
    - action_restart

## interactive_story_1
* greet{"greet": "Heya"}
    - utter_greet
* restaurant_search{"location": "New York"}
    - slot{"location": "New York"}
    - action_chk_location
    - slot{"location_validity": "invalid"}
* goodbye
    - utter_goodbye
    - action_restart


## interactive_story_1
* greet
    - utter_greet
* restaurant_search{"location": "Pittsburgh"}
    - slot{"location": "Pittsburgh"}
    - action_chk_location
    - slot{"location_validity": "invalid"}
* goodbye
    - utter_goodbye
    - action_restart

## interactive_story_1
* greet{"greet": "Hey!"}
    - utter_greet
* restaurant_search{"cuisine": "south indian", "price": "low", "location": "madurai"}
    - slot{"cuisine": "south indian"}
    - slot{"location": "madurai"}
    - slot{"price": "low"}
    - action_chk_location
    - slot{"location_validity": "valid"}
    - action_search_restaurants
    - slot{"response_list": "Kumar Mess 96A, Town Hall Road, Mela Perumal Maistry St. Melavasi Veethi, Periyar, Madurai has been rated 4.4\nAmsavalli Bhavan 42, East Veli Street, Periyar, Madurai has been rated 4.3\nRamana Mess Narayanapuram, Natham Road, Iyer Bungalow Madurai has been rated 4.3\nMurugan Idly Shop West Masi Street, Southern Railway Colony, Madurai has been rated 4.2\nMadurai Mappilai Virundhu Anna Nagar, Sathamangalam, Madurai has been rated 4.2\nMadurai Bun Porotta Aavin Juntion, Opposite PAMC Hospital, Sivagangai Road, Sathamangalam, Madurai has been rated 4.1\nBismi Fried Chicken 79 Pandiyan Velalar ST, Near Kajimar Street Market, Periyar, Madurai has been rated 4.0\nGrill Spot 4/689, Mahalakshmi Nagar 2nd Street, Iyer Bungalow, Madurai has been rated 3.9\nRoyal Mass Meenakshi Nagar, APK Main Road, Villapuram, Madurai has been rated 3.9\nKeeraikadai 7-I Valmiki street, Near Pattu Mahal SS-Colony  Madurai. has been rated 3.9"}
    - utter_email_check
* get_email_ind{"email_ind": "Don't send Email"}
    - slot{"email_ind": "Don't send Email"}
    - utter_goodbye
    - action_restart

## interactive_story_1
* greet{"greet": "Hey!"}
    - utter_greet
* restaurant_search{"cuisine": "south indian", "price": "low", "location": "madurai"}
    - slot{"cuisine": "south indian"}
    - slot{"location": "madurai"}
    - slot{"price": "low"}
    - action_chk_location
    - slot{"location_validity": "valid"}
    - action_search_restaurants
    - slot{"response_list": "Kumar Mess 96A, Town Hall Road, Mela Perumal Maistry St. Melavasi Veethi, Periyar, Madurai has been rated 4.4\nAmsavalli Bhavan 42, East Veli Street, Periyar, Madurai has been rated 4.3\nRamana Mess Narayanapuram, Natham Road, Iyer Bungalow Madurai has been rated 4.3\nMurugan Idly Shop West Masi Street, Southern Railway Colony, Madurai has been rated 4.2\nMadurai Mappilai Virundhu Anna Nagar, Sathamangalam, Madurai has been rated 4.2\nMadurai Bun Porotta Aavin Juntion, Opposite PAMC Hospital, Sivagangai Road, Sathamangalam, Madurai has been rated 4.1\nBismi Fried Chicken 79 Pandiyan Velalar ST, Near Kajimar Street Market, Periyar, Madurai has been rated 4.0\nGrill Spot 4/689, Mahalakshmi Nagar 2nd Street, Iyer Bungalow, Madurai has been rated 3.9\nRoyal Mass Meenakshi Nagar, APK Main Road, Villapuram, Madurai has been rated 3.9\nKeeraikadai 7-I Valmiki street, Near Pattu Mahal SS-Colony  Madurai. has been rated 3.9"}
    - utter_email_check
* get_email_ind{"email_ind": "Don't send Email"}
    - slot{"email_ind": "Don't send Email"}
    - utter_goodbye
    - action_restart

## interactive_story_1
* greet
    - utter_greet
* restaurant_search{"location": "chennai"}
    - slot{"location": "chennai"}
    - action_chk_location
    - slot{"location_validity": "valid"}
    - utter_ask_cuisine
* restaurant_search{"cuisine": "South Indian"}
    - slot{"cuisine": "South Indian"}
    - utter_ask_budget
* restaurant_search{"price": "mid"}
    - slot{"price": "mid"}
    - action_search_restaurants
    - slot{"response_list": "Momo Sa-Khang by Kailash Kitchen 2/247, Perinbavilas Complex, Opposite Chitra Apartments, Choolaimedu, Chennai has been rated 4.6 with average rate for 2 people as 350\nMezze Abhinav Center, Old 19, New 4, Co-operative Colony, Off Chamiers Road, Alwarpet, Chennai has been rated 4.6 with average rate for 2 people as 700\nFusilli Reasons 2/64, Ormes Main Road, Kilpauk, Chennai has been rated 4.5 with average rate for 2 people as 400\nYaa Mohaideen Biryani Old Shop 4/158, Church Road, Opposite Uzhavar Santhai, Pallavaram, Chennai has been rated 4.4 with average rate for 2 people as 600\nWriter's Cafe 98, Peter's Road, Near Philip's Service Centre, Gopalapuram, Chennai has been rated 4.3 with average rate for 2 people as 600"}
    - utter_email_check
* get_email_ind{"email_ind": "Send Email"}
    - slot{"email_ind": "Send Email"}
    - utter_ask_email
* get_email{"email": "venkat.frais@gmail.com"}
    - slot{"email": "venkat.frais@gmail.com"}
    - action_send_email
    - utter_email_confirm
* goodbye
    - utter_goodbye
    - action_restart

## interactive_story_1
* greet
    - utter_greet
* restaurant_search{"location": "madurai"}
    - slot{"location": "madurai"}
    - action_chk_location
    - slot{"location_validity": "valid"}
    - utter_ask_cuisine
* restaurant_search{"cuisine": "South Indian"}
    - slot{"cuisine": "South Indian"}
    - utter_ask_budget
* restaurant_search{"price": "mid"}
    - slot{"price": "mid"}
    - action_search_restaurants
    - slot{"response_list": "Momo Sa-Khang by Kailash Kitchen 2/247, Perinbavilas Complex, Opposite Chitra Apartments, Choolaimedu, Chennai has been rated 4.6 with average rate for 2 people as 350\nMezze Abhinav Center, Old 19, New 4, Co-operative Colony, Off Chamiers Road, Alwarpet, Chennai has been rated 4.6 with average rate for 2 people as 700\nFusilli Reasons 2/64, Ormes Main Road, Kilpauk, Chennai has been rated 4.5 with average rate for 2 people as 400\nYaa Mohaideen Biryani Old Shop 4/158, Church Road, Opposite Uzhavar Santhai, Pallavaram, Chennai has been rated 4.4 with average rate for 2 people as 600\nWriter's Cafe 98, Peter's Road, Near Philip's Service Centre, Gopalapuram, Chennai has been rated 4.3 with average rate for 2 people as 600"}
    - utter_email_check
* get_email_ind{"email_ind": "Send Email"}
    - slot{"email_ind": "Send Email"}
    - utter_ask_email
* get_email{"email": "venkat.frais@gmail.com"}
    - slot{"email": "venkat.frais@gmail.com"}
    - action_send_email
    - utter_email_confirm
* goodbye
    - utter_goodbye
    - action_restart

## interactive_story_1
* greet
    - utter_greet
* restaurant_search{"cuisine": "american", "location": "mumbai"}
    - slot{"cuisine": "american"}
    - slot{"location": "mumbai"}
    - action_chk_location
    - slot{"location_validity": "valid"}
    - utter_ask_budget
* restaurant_search{"price": "high"}
    - slot{"price": "high"}
    - action_search_restaurants
    - slot{"response_list": "Stacks And Racks Shop 1, Ganga Nivas, Opposite Toyota Showroom, Link Road, Malad West, Mumbai has been rated 4.5 with average rate for 2 people as 1000\nThree Chicks and A Bear Lodha Supremus, Survey 453, Senapati Bapat Marg, Railway Colony, Lower Parel, Mumbai has been rated 4.4 with average rate for 2 people as 800\nDoolally Taproom C18-21, Dalia Industrial Estate, Near Fun Republic, Veera Desai Area, Mumbai has been rated 4.3 with average rate for 2 people as 1800\nToit Mathuradas Mill Compound, Senapati Bapat Marg, Lower Parel, Mumbai has been rated 4.2 with average rate for 2 people as 1600\nDoolally Taproom Ground Floor, Raj Kutir, 3rd Road, Near Railway Station, Khar, Mumbai has been rated 4.2 with average rate for 2 people as 1800\nDoolally Taproom 1-3-5, 1st Floor, Kwality House, N.S Patkar Marg, Kemps Corner, Mumbai has been rated 4.1 with average rate for 2 people as 1800\nD:OH! Unit 2A, Ground Floor, Fun Republic, Off New Link Road, Andheri Lokhandwala, Andheri West, Mumbai has been rated 4.0 with average rate for 2 people as 1600\nD:OH! 37A, 1st Floor, Level 1, Viviana Mall, Eastern Express Highway, On Service Road, Near Jupiter Hospital, Majiwada, Thane West, Thane has been rated 4.0 with average rate for 2 people as 1600"}
    - utter_email_check
* get_email_ind{"email_ind": "Send Email"}
    - slot{"email_ind": "Send Email"}
    - utter_ask_email
* get_email{"email": "venkat.frais@gmail.com"}
    - slot{"email": "venkat.frais@gmail.com"}
    - action_send_email
    - utter_email_confirm
* affirm
    - utter_goodbye
    - action_restart

## interactive_story_1
* greet
    - utter_greet
* restaurant_search{"cuisine": "american", "location": "mumbai"}
    - slot{"cuisine": "american"}
    - slot{"location": "mumbai"}
    - action_chk_location
    - slot{"location_validity": "valid"}
    - utter_ask_budget
* restaurant_search{"price": "high"}
    - slot{"price": "high"}
    - action_search_restaurants
    - slot{"response_list": "Stacks And Racks Shop 1, Ganga Nivas, Opposite Toyota Showroom, Link Road, Malad West, Mumbai has been rated 4.5 with average rate for 2 people as 1000\nThree Chicks and A Bear Lodha Supremus, Survey 453, Senapati Bapat Marg, Railway Colony, Lower Parel, Mumbai has been rated 4.4 with average rate for 2 people as 800\nDoolally Taproom C18-21, Dalia Industrial Estate, Near Fun Republic, Veera Desai Area, Mumbai has been rated 4.3 with average rate for 2 people as 1800\nToit Mathuradas Mill Compound, Senapati Bapat Marg, Lower Parel, Mumbai has been rated 4.2 with average rate for 2 people as 1600\nDoolally Taproom Ground Floor, Raj Kutir, 3rd Road, Near Railway Station, Khar, Mumbai has been rated 4.2 with average rate for 2 people as 1800\nDoolally Taproom 1-3-5, 1st Floor, Kwality House, N.S Patkar Marg, Kemps Corner, Mumbai has been rated 4.1 with average rate for 2 people as 1800\nD:OH! Unit 2A, Ground Floor, Fun Republic, Off New Link Road, Andheri Lokhandwala, Andheri West, Mumbai has been rated 4.0 with average rate for 2 people as 1600\nD:OH! 37A, 1st Floor, Level 1, Viviana Mall, Eastern Express Highway, On Service Road, Near Jupiter Hospital, Majiwada, Thane West, Thane has been rated 4.0 with average rate for 2 people as 1600"}
    - utter_email_check
* get_email_ind{"email_ind": "Send Email"}
    - slot{"email_ind": "Send Email"}
    - utter_ask_email
* get_email{"email": "venkat.frais@gmail.com"}
    - slot{"email": "venkat.frais@gmail.com"}
    - action_send_email
    - utter_email_confirm
* affirm
    - utter_goodbye
    - action_restart

## interactive_story_1
* greet
    - utter_greet
* restaurant_search{"cuisine": "american", "location": "mumbai"}
    - slot{"cuisine": "american"}
    - slot{"location": "mumbai"}
    - action_chk_location
    - slot{"location_validity": "valid"}
    - utter_ask_budget
* restaurant_search{"price": "high"}
    - slot{"price": "high"}
    - action_search_restaurants
    - slot{"response_list": "Stacks And Racks Shop 1, Ganga Nivas, Opposite Toyota Showroom, Link Road, Malad West, Mumbai has been rated 4.5 with average rate for 2 people as 1000\nThree Chicks and A Bear Lodha Supremus, Survey 453, Senapati Bapat Marg, Railway Colony, Lower Parel, Mumbai has been rated 4.4 with average rate for 2 people as 800\nDoolally Taproom C18-21, Dalia Industrial Estate, Near Fun Republic, Veera Desai Area, Mumbai has been rated 4.3 with average rate for 2 people as 1800\nToit Mathuradas Mill Compound, Senapati Bapat Marg, Lower Parel, Mumbai has been rated 4.2 with average rate for 2 people as 1600\nDoolally Taproom Ground Floor, Raj Kutir, 3rd Road, Near Railway Station, Khar, Mumbai has been rated 4.2 with average rate for 2 people as 1800\nDoolally Taproom 1-3-5, 1st Floor, Kwality House, N.S Patkar Marg, Kemps Corner, Mumbai has been rated 4.1 with average rate for 2 people as 1800\nD:OH! Unit 2A, Ground Floor, Fun Republic, Off New Link Road, Andheri Lokhandwala, Andheri West, Mumbai has been rated 4.0 with average rate for 2 people as 1600\nD:OH! 37A, 1st Floor, Level 1, Viviana Mall, Eastern Express Highway, On Service Road, Near Jupiter Hospital, Majiwada, Thane West, Thane has been rated 4.0 with average rate for 2 people as 1600"}
    - utter_email_check
* get_email_ind{"email_ind": "Send Email"}
    - slot{"email_ind": "Send Email"}
    - utter_ask_email
* get_email{"email": "venkat.frais@gmail.com"}
    - slot{"email": "venkat.frais@gmail.com"}
    - action_send_email
    - utter_email_confirm
* affirm
    - utter_goodbye
    - action_restart

## interactive_story_1
* greet
    - utter_greet
* restaurant_search{"cuisine": "american", "location": "mumbai"}
    - slot{"cuisine": "american"}
    - slot{"location": "mumbai"}
    - action_chk_location
    - slot{"location_validity": "valid"}
    - utter_ask_budget
* restaurant_search{"price": "high"}
    - slot{"price": "high"}
    - action_search_restaurants
    - slot{"response_list": "Stacks And Racks Shop 1, Ganga Nivas, Opposite Toyota Showroom, Link Road, Malad West, Mumbai has been rated 4.5 with average rate for 2 people as 1000\nThree Chicks and A Bear Lodha Supremus, Survey 453, Senapati Bapat Marg, Railway Colony, Lower Parel, Mumbai has been rated 4.4 with average rate for 2 people as 800\nDoolally Taproom C18-21, Dalia Industrial Estate, Near Fun Republic, Veera Desai Area, Mumbai has been rated 4.3 with average rate for 2 people as 1800\nToit Mathuradas Mill Compound, Senapati Bapat Marg, Lower Parel, Mumbai has been rated 4.2 with average rate for 2 people as 1600\nDoolally Taproom Ground Floor, Raj Kutir, 3rd Road, Near Railway Station, Khar, Mumbai has been rated 4.2 with average rate for 2 people as 1800\nDoolally Taproom 1-3-5, 1st Floor, Kwality House, N.S Patkar Marg, Kemps Corner, Mumbai has been rated 4.1 with average rate for 2 people as 1800\nD:OH! Unit 2A, Ground Floor, Fun Republic, Off New Link Road, Andheri Lokhandwala, Andheri West, Mumbai has been rated 4.0 with average rate for 2 people as 1600\nD:OH! 37A, 1st Floor, Level 1, Viviana Mall, Eastern Express Highway, On Service Road, Near Jupiter Hospital, Majiwada, Thane West, Thane has been rated 4.0 with average rate for 2 people as 1600"}
    - utter_email_check
* get_email_ind{"email_ind": "Send Email"}
    - slot{"email_ind": "Send Email"}
    - utter_ask_email
* get_email{"email": "venkat.frais@gmail.com"}
    - slot{"email": "venkat.frais@gmail.com"}
    - action_send_email
    - utter_email_confirm
* affirm
    - utter_goodbye
    - action_restart

## interactive_story_1
* greet
    - utter_greet
* restaurant_search{"cuisine": "american", "location": "mumbai"}
    - slot{"cuisine": "american"}
    - slot{"location": "mumbai"}
    - action_chk_location
    - slot{"location_validity": "valid"}
    - utter_ask_budget
* restaurant_search{"price": "high"}
    - slot{"price": "high"}
    - action_search_restaurants
    - slot{"response_list": "Stacks And Racks Shop 1, Ganga Nivas, Opposite Toyota Showroom, Link Road, Malad West, Mumbai has been rated 4.5 with average rate for 2 people as 1000\nThree Chicks and A Bear Lodha Supremus, Survey 453, Senapati Bapat Marg, Railway Colony, Lower Parel, Mumbai has been rated 4.4 with average rate for 2 people as 800\nDoolally Taproom C18-21, Dalia Industrial Estate, Near Fun Republic, Veera Desai Area, Mumbai has been rated 4.3 with average rate for 2 people as 1800\nToit Mathuradas Mill Compound, Senapati Bapat Marg, Lower Parel, Mumbai has been rated 4.2 with average rate for 2 people as 1600\nDoolally Taproom Ground Floor, Raj Kutir, 3rd Road, Near Railway Station, Khar, Mumbai has been rated 4.2 with average rate for 2 people as 1800\nDoolally Taproom 1-3-5, 1st Floor, Kwality House, N.S Patkar Marg, Kemps Corner, Mumbai has been rated 4.1 with average rate for 2 people as 1800\nD:OH! Unit 2A, Ground Floor, Fun Republic, Off New Link Road, Andheri Lokhandwala, Andheri West, Mumbai has been rated 4.0 with average rate for 2 people as 1600\nD:OH! 37A, 1st Floor, Level 1, Viviana Mall, Eastern Express Highway, On Service Road, Near Jupiter Hospital, Majiwada, Thane West, Thane has been rated 4.0 with average rate for 2 people as 1600"}
    - utter_email_check
* get_email_ind{"email_ind": "Send Email"}
    - slot{"email_ind": "Send Email"}
    - utter_ask_email
* get_email{"email": "venkat.frais@gmail.com"}
    - slot{"email": "venkat.frais@gmail.com"}
    - action_send_email
    - utter_email_confirm
* affirm
    - utter_goodbye
    - action_restart


## interactive_story_1
* greet
    - utter_greet
* restaurant_search{"cuisine": "chinese", "location": "chennai"}
    - slot{"cuisine": "chinese"}
    - slot{"location": "chennai"}
    - action_chk_location
    - slot{"location_validity": "valid"}
    - utter_ask_budget
* restaurant_search{"price": "mid"}
    - slot{"price": "mid"}
    - action_search_restaurants
    - slot{"response_list": "The Noodle Theory 106/26, Landon's Road, Kilpauk, Chennai has been rated 4.2 with average rate for 2 people as 450\nGreat Wall E-97, 3rd Main Road, Besant Nagar, Chennai has been rated 4.1 with average rate for 2 people as 550\nDragon's Den New 8/2, Old 5/2, MG Road, Adyar, Chennai has been rated 4.0 with average rate for 2 people as 500\nDallay 18/11, Nagendra Nagar, Velachery, Chennai has been rated 4.0 with average rate for 2 people as 350\nHuTong 34, Co-Operative Colony, Alwarpet, Chennai has been rated 3.9 with average rate for 2 people as 700"}
    - utter_email_check
* get_email_ind{"email_ind": "Don't send Email"}
    - slot{"email_ind": "Don't send Email"}
    - utter_goodbye
    - action_restart
