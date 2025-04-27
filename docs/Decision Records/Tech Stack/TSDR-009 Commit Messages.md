# TSDR-009 Commit Messages

## Status

Accepted

## Context

Does every source control commit need a message?

## Decision

No. Not every commit to source control needs a message. For now. If the project
matures and starts to get commits from other developers it will be less and less
appropriate to leave off commit messages.

## Why / Notes

For this decision there are 2 phases in a software projects life cycle. 1st
phase is new green field development. 2nd phase is when a project is mature. 
Multiple developers are contributing to the code base.

In mature projects every commit should have a message. In a mature project every
change should have a separate branch that is tied to a ticket. This is important
for communication between developers.

For the initial green field phase of development, when I'm the solo developer, I've
found that I *personally* don't use commit messages after the fact unless the
commit is substantial. Commits need to be something I would want to revert or
some significant change to code. I remember the old days of frequent computer
crashes. I tend to "save" / commit often. Because of this I tend to have a lot
of small minor commits. For these commits I don't need or even want a message. A
lack of a message is short hand for me that says I can probably ignore that commit.
This works for me and the way I write code for myself.

A commit message should add value. I've found for small changes the file diffs
add more than enough context.

Personally I use source control by the file diffs not the messages. I've learned 
from experience that commit messages lie! Just often enough some change is part
of the commit that isn't covered by the message. These lies *shouldn't* happen
but they do. Because of this by the time I'm reviewing changes I personally don't
usually bother with reading the messages beyond a cursory glance. I have learned
that some kinds of developers *do* rely on and expect them so in a group setting
I still recommend them.

## Consequences

Some developers will think I'm an unprofessional hack. :-)
