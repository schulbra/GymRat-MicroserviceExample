# CRUD.Microservice.Implementation
Portfolio Project for Software Eng I requiring implementation of a microservice.

//----------------------------------------------------------------------------------------------------------------//

- Backend database and interactions for personal project in the works titled " GymRat ".
- Below is the project proposal submitted as an assignment for CS 361 SE I:

 - Project attempted to embody the quality attributes of compatibility, modularity and scalability.
 - The project's software idea is likely to be designed to run via users' cell phone. Although Apple dominates the      market and as such most would be using ios, this idea relies on the ability to easily and effectively collect        user data to be compared publicly and as such I need to accommodate for the various platforms users may have in      such settings. The idea will only be successful if the software is popularized in such settings, so working with    as many individual devices as possible is crucial and potentially very annoying to implement.
 - Modularity was chosen as a means of simplifying the experience of interacting with the software by its users and    the crafting of it via developers . I’d like to collect data to compare that couldn't be altered by a user's        input, or even require it. Something that passively collects and verifies that this comparison of gym activities    isn't falsified. This is the crucial component differentiating it from the saturated fitness tracking app market/    ios’s health feature(s). Ideally the app build could be drafted to contain modules/sections like:
  
   1. Home Gym: The default facility that one works out at and whose name would be on the particular leaderboard.         Its your home base and eventually contains features for interacting with members of the facility you                 participate at.
   2. Users Fitness Data: All personal data making up the criteria displayed via leaderboard(s).User could only view       their own data/stats, but could via other members data via a leaderboard that would be on one of the initial         screens.
   3. Non-User Fitness Data: Non-Personal data that the user would be compared against via leaderboard, or anyone         other than the users workout data.
   4. Leaderboard: Where user data is displayed and compared against specific to the facility they have a membership       at. The leaderboard will consist of a top 25 and reveal a score that is calculated based off of whatever             metrics I decide to use as variables in its equation. Similar to top 25 scores seen in rankings like NCAA           top25.
       a. Individual users would be able to see the individual variables/stats that give rise to this score,                   however, they won't be able to see the components making up the top 25 values of other users. The same               equation would be applied to all users as a means of standardizing data.
   5. To Be Determined Method of Collecting Data + Data options collected: The app needs to be simple via what I           mentioned above, but also so that it could scale to encompass displaying and comparing data of more robust           means, such as entire geographic areas (states, cities, particular franchises, university rec centers, ect). 
        a. Data tracking needs to be started/finished as the user enters/leaves the fitness facility. This data then         needs to be added to the users all-time data total, then compared against current leaderboard standings. If         the metric is >= values in leaderboard, alter the leaderboard to include a new entry. This needs to be done         using whatever device the user is bringing with them to their gym. This could be tested by creating sample           users whom attend a sample home gym all with sample data that is accumulated and compared, however, I don’t         want to type anymore.
 - Simplicity was selected as an attribute primarily because for many working out is already an unpleasent experience. An app that functions as a sort of game might make it more enjoyable but it definitely cant come off as an additional task on top of the other variables that come into play ie: commuting, additional clothes, other people, doms.
