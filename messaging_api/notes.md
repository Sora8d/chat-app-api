- Gonna have to think of a way to invite people to chat groups, although maybe this is a problem for the front (?)
- updateUserConversationLasAccess will remain an endpoint for now, but once jwt headers are available, it should be inside GetMessages. 

- Have to see wether let Front fetch the users_profile or give it to them with GetConversations. Also with this, should then users be moved closer to the messaging api? Get direct access to the table or ...?

-Fix in conversation_dtf doesnt work since ll end up grabbing the same uuid, gotta fix it.