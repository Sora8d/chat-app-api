- Gonna have to think of a way to invite people to chat groups, although maybe this is a problem for the front (?)

- Add constraints

- Later see whether having last msg uuid instead of a timestamp is worth it. Also maybe send less info to the front, like change the proto UserConversation.

BUGS:
- The way Message checks if an user exists makes it able for non participants of the conversation to send messages to conversations. Have to fix it. 
- User-conversations for some reason dont update at certain times. (Error happens in CreateGetMEssage test with flush commented out, after the first iteration.)

- Add a way to insert multiple UserConversations, current one is not efficient for multiple instances.