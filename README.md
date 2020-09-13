# eventBusGeneric

Supports  drivers such as InMemory / Redis for Pub-Sub Event based model.

- InMemory Event Bus implementation of the pub/sub pattern where publishers are publishing data and <br> interested subscribers can listen to them and act based on data. InMemory Event Bus uses channels, <br> Subscribers subscribes to particular topic and publisher publishes data on topic.
The channel will <br> receive the data when publisher publishes data to the topic

- Redis Pub/Sub messaging paradigm allows applications talk to each other through subscription to channels.<br>
This allows decoupling between publishers and subscribers, Subscribers express interest in one or more <br> channels, and only receive messages that are of interest,  without knowledge of any publishers.