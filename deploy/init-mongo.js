const db = new Mongo().getDB('search_engine');
db.conversation.drop();
db.createCollection("conversation")
db.conversation.createIndexes([
    { "prompt": "text", "response": "text"},
    { "id": 1 },
]);
db.conversation.insertMany([
    {
        id: 1,
        prompt: "What is your name ?",
        response: "My name is jarvis ."
    },
    {
        id: 2,
        prompt: "What is your age ?",
        response: "I am 1 year old ."
    },
    {
        id: 3,
        prompt: "What is your sex ?",
        response: "I am a bot ."
    }
]);