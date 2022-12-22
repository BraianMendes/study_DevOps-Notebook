# Best Practices to Optimize Infrastructure Monitoring within DevOps Teams

It’s no secret that the lines between “dev” and “ops” have blurred. Traditionally, a system administrator (or sysadmin) is responsible for maintaining reliable operations, such as the upkeep of performance and security. On the other hand, developers are responsible for building, coding and deploying software.

Now, these two roles have virtually become one as enterprise IT environments have become increasingly complex. To manage growing infrastructure requirements, development and operations teams learned that they can no longer work in siloed environments, which led to the modern-day DevOps evolution. Still, both job functions are critical for business continuity, and if organizations do not empower the needs of developers and operators, they will be at a huge disadvantage.

The coming together of DevOps, and its evolving role, has less to do with hiring the developer with the broadest range of skills or adding more team members to solve incoming challenges and more to do with bringing current teams together. However, even though many companies have already made this cultural shift, developers continue to be bogged down by cumbersome systems that rob them of their time and creative freedom. To bypass these bottlenecks and potential burnout there are a few best practices that DevOps teams can implement to make infrastructure management easier and more intuitive.

It all boils down to a DevOps team’s monitoring capabilities, which can provide actionable insights for even the most technical team members, including those who may not traditionally have needed to monitor the health of an organization’s IT infrastructure before.

## Granular Visibility Matters

Infrastructure monitoring, when implemented correctly, can be segmented into two broad categories: reactive and proactive. Both are required to ensure systems operate at peak efficiency, which is only achieved when all available computer resources are consumed, monitored and analyzed so that business goals can come to fruition. These goals can range from the availability or uptime of a website to the response time of a particular feature. The reason DevOps teams should aim for peak efficiency is so that optimal resource allocation can be identified. If teams over-procure computer resources, they waste unused computational power, which increases costs needlessly. On the other hand, if teams under-procure, then there is potential for business goals to not be achieved and accrue a monetary loss. For example, research has shown that for every second of downtime, a certain percentage of users will leave a website.

Reactive infrastructure monitoring is all about measuring and tracking its most quintessential aspects — those that run the application — and studying the relationship between the application and the underlying ecosystem to respond to issues as they occur. This may be tedious work but is essential to ensuring critical business outcomes.

Proactive infrastructure monitoring is the type of monitoring that DevOps professionals would prefer to spend their time on. Since reactive monitoring focuses on system incidents after they occur, proactive monitoring enables DevOps professionals to find the root cause of an issue before it occurs. By anticipating or monitoring for past issues, teams can address them before they jeopardize production systems even further.

Key infrastructure metrics that DevOps teams should monitor include Random Access Memory (RAM), Central Processing Units (CPUs), the network, and storage, which are all critical for an application’s proper behavior. By closely monitoring these aspects, developers can better understand the most basic behavior of their application(s) and decide on the infrastructure required to run it, since costs can easily spiral out of control if the developer is not paying attention.

Similarly, it is not uncommon for developers to suddenly realize that their application is not performing as planned or is consuming more resources than anticipated due to a design flaw. This translates into heightened costs, particularly with elastic deployments, where the developer is not buying specific Virtual Machines (VMs) with specific hardware attributes but rather abstractions that scale automatically.

Brendan Gregg, a very prominent performance engineer, outlines many ways to get core metrics from Linux systems to satisfy a framework he calls the USE method for performance analysis (utilization, saturation, and errors). The more perceptive DevOps professional will note that the USE method shares common attributes with the four golden signals of monitoring, namely Latency, Traffic, Errors and Saturation, which were created by Google’s site reliability engineers (SREs) and were introduced in the highly influential Google SRE book.

For example, to measure CPUs, developers could use the following signals:

<ul>
<li>Latency: Delay in CPU scheduling</li>
<li>Traffic: CPU utilization</li>
<li>Errors: CPU-related errors (e.g. number of faulted CPUs)</li>
<li>Saturation: Length of the run queue</li>
</ul>


Regardless of the metrics that your team chooses to use for reactive and proactive infrastructure monitoring, the key is to equip DevOps teams with true insight into their application’s performance, as granular as possible, with data to back up their findings.

## Making Monitoring Seamless

More often than not, especially in the early iterations of a product, the development team oversees the bulk of operations (ironic, right?). Although they may be assisted by a DevOps or SRE professional when setting up the infrastructure or deciding on certain deployment options, developers are ultimately responsible for monitoring their application and ensuring that it meets business goals. To that effect, simple-to-use tools can boost productivity since the developer won’t need to focus on the complexity of the monitoring tool but rather the complexity of the application. If you are a developer yourself, you may be familiar with the pull to spend as little time as possible setting up a monitoring tool (it’s alarms and functions) so that you can do what you want — writing code and developing your application’s features.

Although it may be tempting to utilize a zero-configuration monitoring tool, developers should be aware that their systems should be highly customizable and provide easy integration into the toolchain. Finding tools that integrate well into existing processes can make a developer’s job just easier. In that respect, a continuous integration/continuous delivery (CI/CD) pipeline that includes stress and load testing on the application can be highly beneficial. By installing a monitoring solution as part of the CI/CD, the user can create alarms to signal to the CI/CD pipeline if a test fails. For example, an alarm can be triggered if the RAM consumption goes over a particular amount.

In general, there is a move toward abstraction and a general understanding that the user does not wish to set up everything within a monitoring tool themselves. Ease of use and sound UX/UI design have returned as major players in the monitoring industry to embrace these principles. The infrastructure tools that exist today are much more user-friendly than they were 20 years ago, meaning that more junior DevOps or developer professionals can use them. Monitoring an entire stack is now not strictly a senior-level skill. New tools on the market not only have simple defaults, but they require very little configuration and can provide comprehensive insights about infrastructure systems. In this vein, the more opinionated a system is, the more decisions are lifted from the user and are taken up by the product itself.

To ensure DevOps teams are able to meet business goals, implementing monitoring tools that are easy to use and setup — and require limited time for maintenance — have an edge and allow both devs and ops to be the same. They also ensure that DevOps teams have a granular view of the infrastructure in play, and can act reactively and proactively as the situation requires.