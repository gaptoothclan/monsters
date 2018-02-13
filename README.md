# Monsters Coding Challenge

Wrote the original solution in node.js, in a very simple array and object loop until the end style. I found this to be a very good solution to the problem, but code wise I was not very happy. Some of my concerns included, not being at all testable, not being very readable, performance wise I felt it could be better by managing the relationships between the nodes and handling the responsibilities at an object level, also this test really deals with state so a more object orientated design seemed to fit.

The code I have written here far exceeds the 2 hours for this challenge, albeit sat infront of the TV with baby duties distracting me. The general hieracy arcitecture of this is as follows:

Main
  World 
    City
      Monster
      
Cities also know about the cities that they are linked to so they can notify each other when they die.

Overall I would feel reasonably happy about putting this into a live enviroment.

