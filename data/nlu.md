## intent:affirm
- yes
- yep
- yeah
- indeed
- that's right
- ok
- great
- right, thank you
- correct
- great choice
- sounds really good
- thanks
- cool
- awesome
- thank you
- Yes
- Great. Bye!
- Thank you
- Ok Cool
- Sure
- hmm, i'd like that
- sure!
- yes i accept
- Sweet
- amazing!
- how nice!
- cool!
- yay
- yes accept please
- great
- oh cool
- yes
- fine
- i will take that
- that sounds just right

## intent:get_email_ind
- [Send Email](email_ind)
- [Don't send Email](email_ind)

## intent:goodbye
- bye
- goodbye
- good bye
- stop
- end
- farewell
- Bye bye
- have a good one
- Alryt
- All right
- Bye

## intent:greet
- hey
- howdy
- hey there
- hello
- hi
- good morning
- good evening
- dear sir
- Hola
- Hiya
- Heya
- Hey
- hey dude
- hello it is me again
- what up
- hi there
- good evening
- good morning
- good afternoon

## intent:restaurant_search
- i'm looking for a place to eat
- I want to grab lunch
- I am searching for a dinner spot
- I am looking for some restaurants in [Delhi](location).
- I am looking for some restaurants in [Bangalore](location)
- show me [chinese](cuisine) restaurants
- show me [chines]{"entity": "cuisine", "value": "chinese"} restaurants in the [New Delhi]{"entity": "location", "value": "Delhi"}
- show me a [mexican](cuisine) place in the [centre](location)
- i am looking for an [indian](cuisine) spot called olaolaolaolaolaola
- search for restaurants
- anywhere in the [west](location)
- I am looking for [asian fusion](cuisine) food
- I am looking a restaurant in [294328](location)
- in [Gurgaon](location)
- [South Indian](cuisine)
- [North Indian](cuisine)
- [Italian](cuisine)
- [Chinese]{"entity": "cuisine", "value": "chinese"}
- [mid](price)
- [low](price)
- [high](price)
- [chinese](cuisine)
- [Lithuania](location)
- Oh, sorry, in [Italy](location)
- in [delhi](location)
- I am looking for some restaurants in [Mumbai](location)
- I am looking for [mexican indian fusion](cuisine)
- can you book a table in [rome](location) in a [moderate]{"entity": "price", "value": "mid"} price range with [british](cuisine) food for [four]{"entity": "people", "value": "4"} people
- [central](location) [indian](cuisine) restaurant
- please help me to find restaurants in [pune](location)
- Please find me a restaurantin [bangalore](location)
- [mumbai](location)
- show me restaurants
- please find me [chinese](cuisine) restaurant in [delhi](location)
- can you find me a [chinese](cuisine) restaurant
- [delhi](location)
- please find me a restaurant in [ahmedabad](location)
- please show me a few [italian](cuisine) restaurants in [bangalore](location)
- can you find me a restaurant
- [chennai](location)
- Can you find me a restaurant in [merut](location)
- can you find me a restaurant in [madurai](location)
- can you find me a restaurant in [Chengalpat](location)
- can you find me a restaurant in [chengalpat](location)
- can you find me a restaurant in [chennai](location)
- can you find me a restaurant in [wayanad](location)
- Can you find me a restaurant
- Can you find me a restaurant in [Chennai](location)?
- Can you find me a [south indian](cuisine) restaurant in [madurai](location)?
- Can you find me a [south indian](cuisine) restaurant in [chennai](location)?
- I'm Hungry. Looking for a restaurant
- [Madurai](location)
- Can you find me a [cafe](cuisine) in [chennai](location)?
- can you find me a [high](price) range [north indian](cuisine) restaurant in [mumbai](location)?
- can you find me a [low](price) range [north indian](cuisine) restaurant in [mumbai](location)?
- can you find me a [mid](price) range [north indian](cuisine) restaurant in [mumbai](location)?
- can you find me a [high](price) range restaurant in [delhi](location)?
- can you find me a [mid](price) range restaurant in [delhi](location)?
- Can you suggest some good restaurants in [Faridabad](location)?
- Can you find me a [mexican](cuisine) restaurant in [mumbai](location)?
- Can you find me a restaurant in [New York](location)?
- Can you find me a restaurant in [Pittsburgh](location)?
- Can you find a [south indian](cuisine) restaurant with [low](price) price in [madurai](location)?
- Can you find a restaurant in [chennai](location)
- Can you find an [american](cuisine) restaurant in [mumbai](location)?
- Can you find me a [chinese](cuisine) restaurant in [chennai](location)?
- [mid](price)

## intent:get_email
- [Yes]{"entity": "email", "value": "anand19.r@gmail.com"}
- [anand19.r@gmail.com](email)
- [venkat.frais@gmail.com](email)

## synonym:4
- four

## synonym:Delhi
- New Delhi
- dilli
- newdelhi
- Newdelhi
- new delhi
- new Delhi

## synonym:anand19.r@gmail.com
- Yes

## synonym:bangalore
- Bengaluru
- bengaluru
- banglore

## synonym:north indian
- north-indian
- northindian
- north-indina
- North Indian

## synonym:south indian
- south-indian
- southindian
- south-indina
- South Indian

## synonym:chennai
- madras
- Chennai

## synonym:chinese
- chines
- Chinese
- Chines

## synonym:kolkata
- calcutta
- kolkatta
- calcuta

## synonym:mumbai
- Bombay
- mumbai
- bombay

## synonym:mysore
- Mysore
- mysuru
- Mysuru

## synonym:kochi
- kochi
- cochin
- Cochin

## synonym:Thiruvananthapuram
- thiruvananthapuram
- Trivandrum
- Travancore

## synonym:Pondicherry
- puducherry
- pondy

## synonym:mid
- moderate

## synonym:vegetarian
- veggie
- vegg

## regex:greet
- hey[^\s]*

## regex:email
- ^[a-z0-9]+[\._]?[a-z0-9]+[@]\w+[.]\w{2,3}$

## regex:pincode
- [0-9]{6}
