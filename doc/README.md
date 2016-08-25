# Disclaimer

DISCLAIMER: I'm not expecting anyone to look at this project, or use it, or contribute to it, for
quite a while. I'm using it as a learning exercise more than anything else at the moment, with the
goal of eventually knowing go, bootstrap, backbone, and so on. I will remove this disclaimer when
I think the project has reached MVP (minimum viable project) stage. That won't be for a while.

# funtodone
Adding a little random fun (and a little extra intelligence) to your todo list

# The Enemy

Distraction, procrastination, and constant interruptions are killing our productivity, our self-esteem, our ability
to concentrate, and (possibly due to increased stress as a result) even our health. 

But lets face it -- we're knowledge workers. We can't go hide in the forest and cut ourselves off from the Internet
and still get our work done. 

If you look up procrastination in Amazon, you'll currently find 7129 results among books.
A search for distraction returns another 1775. There are LOTS of software tools already available
that are specifically intended to fight those.

This is another one. But this one is different:

 1. We assume the _*reason*_ people procrastinate is avoidance, not lack of organization
 2. We focus on the _*process*_, and try to make the process itself fun
 3. We add a spoonful of sugar (just like Mary Poppins, back before sugar was evil)

# Why Now?

Things are only going to get worse. 

  * the main economic force driving the internet is advertising
  * advertisers want your eyeballs
  * there's an advertiser standing behind you right now, holding a scalpel!
  * just kidding
  * but seriously, there is more fun and interesting stuff on the net every day: we need to make your work more fun and interesting just to keep up

# Just Imagine

Wouldn't it be great to go to bed at night, and feel like you'd accomplished a lot during the day?

Wouldn't it be great to have visible success, and accomplish big things over time? Imagine that you finally learn
that second language, write that novel, build that dream house, lose that weight, run that marathon,
get that promotion, or create that app everyone's talking about? Or maybe all of the above?

Wouldn't it be great to have a reputation as somebody who gets things done?

# But Wait, What About...?

## Isn't this just some cheesy "gamification" of a ToDo app?

Heck no. You don't need no steenking badges. 

On the other hand, one of the critical steps in improving anything is to measure it, which is basically a way of keeping score.

And on the third hand, the reason gamification got such a bad name so quickly is that legions of wild-eyed hyper-ventilating consultants
oversold it to such a degree that it became a joke before anyone even had a chance to see if there was anything to it.

It *is* possible to make the process of planning and keeping track of your actions more fun than you might expect, even if it's
unlikely to make you laugh out loud with glee.

## I don't want to spend hours (or days) entering all my personal data into somebody else's database

Yeah, neither do we. That's why we don't store your data ourselves. Fortunately, there are lots of little companies like
Google (Drive), Microsoft (OneDrive), and Apple (iCloud) -- not to mention Dropbox, Box, and others -- that let you save
your data in cloud storage that is shared across all your devices. That's what we're going to try to use.

We store all your raw data in JSON format (which means it's easy to import and export, and you can even read it yourself if
you tilt your head and squinch your eyes a bit).

## Where did the ideas for this come from?

It's that Poppins woman! She ... well, that *is* one source, I suppose. But more seriously, here are some of the
sources for the ideas behind funtodone:

 * [David Allen's "Getting Things Done"](http://gettingthingsdone.com/)
 * [Daniel Pink's "Drive: The Surprising Truth About What Motivates Us"](http://www.amazon.com/Drive-Surprising-Truth-About-Motivates/dp/1594484805)
 * [Dan Arielly's "Predictably Irrational: The Hidden Forces That Shape Our Decisions"](http://www.amazon.com/Predictably-Irrational-Revised-Expanded-Edition/dp/0061353248)
 * [Carol Dweck's "Mindset: The New Psychology of Success"](http://www.amazon.com/Mindset-Psychology-Success-Carol-Dweck/dp/0345472322)
 * [Kelly McGonigal's "The Willpower Instinct: How Self-Control Works, Why It Matters, and What You Can Do to Get More of It"](http://www.amazon.com/Willpower-Instinct-Self-Control-Works-Matters-ebook/dp/B005ERIRZE)
 * [Jane McGonigal's "Reality Is Broken: Why Games Make Us Better and How They Can Change the World"](http://www.amazon.com/Reality-Broken-Games-Better-Change-ebook/dp/B004G8Q1Q4)

## Where are all the testimonials?

Shouldn't we have all kinds of glowing testimonials from celebrities and People With Credentials, saying how
wonderful _funtodone_ is, and that it helped them immeasurably and made them who they are today?

Yeah, I guess. We probably ought to wait until we have something working, though, and then see about getting somebody to try it out.

## This is going to be too complicated to use, isn't it?

Nope! ...although you can fiddle with options to make it more complicated once you get the hang of it.

But here's how you play, in a nutshell -- in general, three clicks and you're doing something:

  1. Click a button to get a task or make a new one:
    a. new task
    b. hit me! (a task is selected for you, see gory details about how that works below)
    c. choose task (you pick which task)
  2. Once you've got a task, you do one of these:
    a. back to main (to get or create a new task)
    b. edit it
    c. split it (make it into two or more separate tasks)
    d. expand it (add one or more child tasks)
    e. link it (edit what it depends on, or what depends on it)
    f. do it
  3. One you're doing a task, you can:
    a. go back to the previous stage
    b. start a timer
    c. track your progress
    d. edit it
    e. finish it
    f. stop working on it

There's also some configuration options you can play with.

Part of the fun is to let funtodone pick a next task for you to work on. That way if you've got 
some tasks that are urgent but less fun than others, you can either avoid them if you're lucky,
or get them out of the way if you're not. Since we recommend you not work on the same task for
more than 20 minutes at a time, even the least fun tasks aren't too bad.

Gory Details:

Choosing a random task isn't quite all random, since funtodone tries to weight the choices a little:
  * prefer tasks with an imminent due date
  * try not to pick the same task twice in a row (unless it's important and due right away)
  * try to alternate between more fun and less fun tasks
  * try to work on the more important things early in the day

## Too Many Tasks?

You might think that continuously splitting tasks into smaller and smaller pieces will make it
seem overwhelming. 

The trick is that (normally) you're not looking at the whole list, but just a single sub-task or
a parent task and a list of its subtasks.

You want to get the level of difficulty for a task down to an almost rediculously low level,
in order to bypass the resistance that automatically shows up for larger tasks. It's like
cleaning up after a big party -- it can feel overwhelming at first, but if you tell yourself
"OK, I'm just going to spend five minutes washing dishes" then, after that, you look around
and it's not so bad any more.

The Japanese principal of Kaizen works a lot like that. You make small changes rather than 
trying for big heroic ones, and before you know it, really big things get done -- without
all the wailing and teeth gnashing that a large change would entail.

# License

This is open source software, licensed under the MIT license. That seemed like that most permissive
version available. Please see the LICENSE file for details.

