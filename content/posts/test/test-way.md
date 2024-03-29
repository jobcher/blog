---
title: "测试策略：微服务架构"
date: 2023-12-15
draft: true
featuredImage: "/images/test.jpeg"
featuredImagePreview: "/images/test.jpeg"
images: ['/images/test.jpeg']
author: "jobcher"
tags: ["测试"]
categories: ["测试"]
series: ["测试"]
---
<p data-nodeid="886">可以从两个方面来保障微服务的质量：</p>
<ol data-nodeid="887">
<li data-nodeid="888">
<p data-nodeid="889">选取合适的测试策略模型，确保测试活动全面且有效；</p>
</li>
<li data-nodeid="890">
<p data-nodeid="891">建立质量保障体系，使质量保障内化为企业的组织能力。</p>
</li>
</ol>
<h3 data-nodeid="892">如何选择合适的测试策略模型？</h3>
<p data-nodeid="893">要想使面向微服务的测试活动全面且有效，可以借用“测试金字塔”的思想，针对不同类型和颗粒度的测试投入不同的精力，达到一个最佳平衡：</p>
<ul data-nodeid="894">
<li data-nodeid="895">
<p data-nodeid="896">测试需要分层，每一层的测试颗粒度有所不同；</p>
</li>
<li data-nodeid="897">
<p data-nodeid="898">不同层次的测试比重有差异，通常来说，层次越高，测试比重应越少。</p>
</li>
</ul>
<p data-nodeid="899"><img src="https://s0.lgstatic.com/i/image/M00/37/4A/CgqCHl8ZQp2AA2yKAADyJvMVUks187.png" alt="Drawing 0.png" data-nodeid="969"></p>
<p data-nodeid="900">需要说明的是，传统意义下的测试金字塔，在微服务架构下不再完全奏效。因为微服务中最大的复杂性不在于服务本身，而在于微服务之间的交互方式，这一点值得特别注意。</p>
<p data-nodeid="901">因此，<strong data-nodeid="976">针对微服务架构，常见的测试策略模型</strong>有如下几种。</p>
<p data-nodeid="902">（1） <strong data-nodeid="981">微服务“测试金字塔”</strong></p>
<p data-nodeid="1106" class="">基于微服务架构的特点和测试金字塔的原理，Toby Clemson 有一篇关于<a href="https://www.martinfowler.com/articles/microservice-testing/" data-nodeid="1110">“微服务架构下的测试策略”</a>的文章，其中通过分析阐述了微服务架构下的通用测试策略。</p>

<p data-nodeid="904"><img src="https://s0.lgstatic.com/i/image/M00/37/3F/Ciqc1F8ZQrSACTc9AAB65lA45vc729.png" alt="Drawing 1.png" data-nodeid="985"></p>
<p data-nodeid="905">如图，该策略模型依然是金字塔形状，从下到上依次为单元测试、集成测试、组件测试、端到端测试、探索式测试。</p>
<p data-nodeid="906">（2） <strong data-nodeid="991">微服务“测试蜂巢”</strong></p>
<p data-nodeid="907">这种策略模型是蜂巢形状，它强调重点关注服务间的集成测试，而单元测试和端到端测试的占比较少。</p>
<p data-nodeid="908"><img src="https://s0.lgstatic.com/i/image/M00/37/4A/CgqCHl8ZQsGAZti7AABGRbBNFY8164.png" alt="Drawing 3.png" data-nodeid="995"></p>
<p data-nodeid="909">（3） <strong data-nodeid="1000">微服务“测试钻石”</strong></p>
<p data-nodeid="910">这种策略模型是钻石形状的，组件测试和契约测试是重点，单元测试比率减少，另外增加了安全和性能等非功能的测试类型。</p>
<p data-nodeid="911"><img src="https://s0.lgstatic.com/i/image/M00/37/4A/CgqCHl8ZQs-AByNAAACgJaZwyyU241.png" alt="Drawing 5.png" data-nodeid="1004"></p>
<p data-nodeid="912">我想，有多少个基于微服务架构的测试团队大概就有多少个测试策略模型吧。<strong data-nodeid="1010">“测试金字塔”是一种测试策略模型和抽象框架</strong>，当技术架构、系统特点、质量痛点、团队阶段不同时，每种测试的比例也不尽相同，而且最关键的，并不一定必须是金字塔结构。</p>
<p data-nodeid="913">理解了测试策略模型的思考框架，我们看下应如何保障测试活动的全面性和有效性。</p>
<h4 data-nodeid="914">全面性</h4>
<p data-nodeid="915">微服务架构下，既需要保障各服务内部每个模块的完整性，又需要关注模块间、服务间的交互。只有这样才能提升测试覆盖率和全面性，因此，可以通过如下的分层测试来保证微服务的全面性。</p>
<p data-nodeid="916"><img src="https://s0.lgstatic.com/i/image/M00/37/51/CgqCHl8ZSrqAVjqcAAVCHyjoRMg887.png" alt="Drawing 7.png" data-nodeid="1016"></p>
<ul data-nodeid="917">
<li data-nodeid="918">
<p data-nodeid="919">单元测试（Unit Test） ：从服务中最小可测试单元视角验证代码行为符合预期，以便测试出方法、类级别的缺陷。</p>
</li>
<li data-nodeid="920">
<p data-nodeid="921">集成测试（Integration Test）：验证当前服务与外部模块之间的通信方式或者交互符合预期，以便测试出接口缺陷。</p>
</li>
<li data-nodeid="922">
<p data-nodeid="923">组件测试 （Component Test）：将测试范围限制在被测系统的一部分（一般是单个服务），使用测试替身（test doubles）将其与其他组件隔离，以便测试出被测代码的缺陷。</p>
</li>
<li data-nodeid="924">
<p data-nodeid="925">契约测试（Contracts Test）：验证当前服务与外部服务之间的交互，以表明它符合消费者服务所期望的契约。</p>
</li>
<li data-nodeid="926">
<p data-nodeid="927">端到端测试（End-to-end Test）：从用户视角验证整个系统的功能能否符合用户的预期。</p>
</li>
</ul>
<p data-nodeid="928">可见，上述测试策略模型中的测试方法，是自下而上逐层扩大测试范围和边界，力保微服务架构的模块内、模块间交互、服务内、服务间交互、系统范围等维度的功能符合预期。</p>
<h4 data-nodeid="929">有效性</h4>
<p data-nodeid="930">确定了分层测试方法，我们应该如何选取每种测试方法的比例，来确保该测试策略的有效性呢？</p>
<p data-nodeid="931">首先必须要明确的是不存在普适性的测试组合比例。我们都知道，测试的目的是解决企业的质量痛点、交付高质量的软件。因此不能为了测试而测试，更不能为了质量而不惜一切代价，需要<strong data-nodeid="1029">考虑资源的投入产出比。</strong></p>
<p data-nodeid="932">测试策略如同测试技术、技术架构一样，并不是一成不变，它会随着业务或项目所处的阶段，以及基于此的其他影响因素的变化而不断演进。但归根结底，还是要从质量保障的目标出发，制定出适合当时的测试策略，并阶段性地对策略进行评估和度量，进而不断改进和优化测试策略。因此，<strong data-nodeid="1034">选取测试策略一定要基于现实情况的痛点出发，结果导向，通过调整测试策略来解决痛点。</strong></p>
<p data-nodeid="933">比如，在项目早期阶段或某 MVP 项目中，业务的诉求是尽快发布到线上，对功能的质量要求不太高，但对发布的时间节点要求非常严格。那这种情况下快速地用端到端这种能模拟用户真实价值的测试方法保障项目质量也未尝不可；随着项目逐渐趋于平稳后，时间要求渐渐有了节奏，对功能的质量要求会逐渐变高，那么这时候可以再根据实际情况引入其他测试方法，如契约测试或组件测试等。</p>
<p data-nodeid="934">你要永远记住，<strong data-nodeid="1040">适合自身项目阶段和团队的测试策略才是“完美”的策略。</strong></p>
<p data-nodeid="935"><img src="https://s0.lgstatic.com/i/image/M00/37/51/CgqCHl8ZSvOAK06pAAVCHyjoRMg396.png" alt="Drawing 7.png" data-nodeid="1043"></p>
<h3 data-nodeid="936">如何建立质量保障体系？</h3>
<p data-nodeid="937">上述分层的测试策略只是尽可能地对微服务进行全面的测试，确保系统的所有层次都被覆盖到，它更多体现在测试活动本身的全面性和有效性方面。要想将质量保障内化为企业的组织能力，就需要通过技术和管理手段形成系统化、标准化和规范化的机制，这就需要建设质量保障体系。</p>
<p data-nodeid="938"><strong data-nodeid="1050">质量保障体系：通过一定的流程规范、测试技术和方法，借助于持续集成/持续交付等技术把质量保障活动有效组合，进而形成系统化、标准化和规范化的保障体系。</strong> 同时，还需要相应的度量、运营手段以及组织能力的保障。</p>
<p data-nodeid="939">如下几个方面是质量保障体系的关键，后续课程也将按如下方式展开讲解。</p>
<ul data-nodeid="940">
<li data-nodeid="941">
<p data-nodeid="942"><strong data-nodeid="1056">流程规范</strong>：没有规矩不成方圆，好的流程规范是保障质量中非常关键的一环。当出现交付质量差的情况时，过程质量也一定是差的。通常会出现某些关键动作未执行或执行不到位、对事情的不当处理等情况，而这些情况可以通过建立闭环、分工明确的流程规范完全可以避免。另外，研发过程中，过程质量跟执行人的质量意识、个人能力等直接相关，那么就需要建立易执行的流程规范，降低人员的执行门槛。同时需要特别注意，规范的不断完善是几乎所有团队的常态，但当规范执行效果不好时一定要及时跟进，分析其根本原因，必要时要进行简化。</p>
</li>
<li data-nodeid="943">
<p data-nodeid="944"><strong data-nodeid="1061">测试技术</strong>： 测试策略模型中的分层测试方法可以使面向微服务的测试活动具有一定的全面性和有效性，使得被测内容在功能方面符合预期。除功能性之外，软件质量还有其他很多属性，如可靠性、易用性、可维护性、可移植性等，而这些质量属性就需要通过各种专项测试技术来保障了。同时，还有许多测试技术的首要价值在于提升测试效率。因此合理地组合这些测试技术，形成测试技术矩阵，有利于最大化发挥它们的价值。</p>
</li>
<li data-nodeid="945">
<p data-nodeid="946"><strong data-nodeid="1066">持续集成与持续交付</strong>：微服务的优势需要通过持续集成和持续交付的支持才能充分发挥出来，这就要求在执行测试活动时提高反馈效率、尽快发现问题。一方面要明确各种“类生产环境”在交付流程中的位置和用途差异点，保证它们的稳定可用。另一方面需要将各种测试技术和自动化技术集成起来，使代码提交后能够自动进行构建、部署和测试等操作，形成工具链，实现真正意义上的持续集成和持续交付。</p>
</li>
<li data-nodeid="947">
<p data-nodeid="948"><strong data-nodeid="1075">度量与运营</strong>：管理学大师德鲁克曾经说过“你如果无法度量它，就无法管理它（It you can’t measure it, you can’t manage it)”。要想能有效管理和改进，就难以绕开度量这个话题。对于研发过程来说，度量无疑是比较难的，它是一个脑力密集型的过程，指标多维度，且很多维度的内容难以清晰地度量。在质量保障体系中，我将基于质量、效率、价值等多维度视角建立起基础的度量体系，并结合定期运营做定向改进，形成 PDCA 正向循环，促使各项指标稳步提升。同时，需要特别警惕的是，<strong data-nodeid="1076">度量是一把双刃剑</strong>，这里我也会告诉一些我的经验教训和踩过的坑，避免走错方向。</p>
</li>
<li data-nodeid="949">
<p data-nodeid="950"><strong data-nodeid="1081">组织保障</strong>：产品的交付离不开组织中每个参与部门成员的努力。正如质量大师戴明所说，质量是设计出来的，不是测试出来的。因此在组织中树立起“质量文化”至关重要。在这部分内容里，我将介绍常见的参与方的角色、职责和协作过程中的常见问题、对策，以及如何营造质量文化等内容。</p>
</li>
</ul>
<h3 data-nodeid="951">总结</h3>
<p data-nodeid="952">谈到了基于微服务架构下的各种质量挑战，可以从两个方面有效且高效地保障微服务的质量：<strong data-nodeid="1091">确保面向微服务的测试活动具备全面性和有效性</strong>，<strong data-nodeid="1092">质量保障需要内化为企业的组织能力。</strong></p>
<p data-nodeid="953">通过对测试金字塔原理和微服务的特点分析，引入单元测试、集成测试、组件测试、契约测试和端到端测试等分层测试类型来确保测试活动的全面性，通过自身项目阶段和团队情况来选取合适的测试策略模型，以保障测试活动的有效性。</p>
<p data-nodeid="954">要想把质量保障内化为企业的组织能力，就需要通过系统的技术和管理手段形成机制，在流程规范、测试技术、持续集成与持续交付、度量与运营、组织保障等方面构建质量保障体系。</p>
<p data-nodeid="955">你是否测试过微服务架构的项目和服务？如果有，欢迎在留言区评论，说说你所经历过的项目的测试策略和质量保障体系是怎样的，期间遇到了哪些困难和问题。同时欢迎你能把这篇文章分享给你的同学、朋友和同事，大家一起交流。</p>
<p data-nodeid="956"><strong data-nodeid="1099">相关链接</strong></p>
<blockquote data-nodeid="957">
<p data-nodeid="958" class="">Testing Strategies in a Microservice Architecture（微服务架构下的测试策略）：<br>
：<a href="https://www.martinfowler.com/articles/microservice-testing/" data-nodeid="1105">https://www.martinfowler.com/articles/microservice-testing/</a></p>
</blockquote>