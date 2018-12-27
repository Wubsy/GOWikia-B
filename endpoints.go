package GOWikia_Wubsy

//Here I will attempt to make it as easy as possible to find what you're looking for.
//I'm going to be basing most of this library around bwmarrin's repo: `discordgo`

var(
	EndpointBaseWikia = "wikia.com/" // Require the user of the lib to supply the subdomain
	EndpointTestWikia = "http://muppet.wikia.com/" //Format follows a slash at the end, rather than having both slashes in a single string
	EndpointAPI = "http://www." + EndpointBaseWikia + "api/"
	EndpointVersion = EndpointAPI + "v1/" //Presently, the only version is v1

	EndpointActivity = EndpointVersion + "activity/"
	EndpointLatestActivity = EndpointActivity + "LatestActivity/"
	EndpointRecentChange = EndpointActivity + "RecentlyChangedArticles/"

	EndpointArticles = EndpointVersion + "Articles/"
	EndpointArticlesAsSimpleJson = EndpointArticles + "AsSimpleJson/"
	EndpointArticlesDetails = EndpointArticles + "Details/"
	EndpointArticlesList = EndpointArticles + "List/"
	EndpointArticlesListAll = EndpointArticles + "List?expand=1"
	EndpointArticlesMostLinked = EndpointArticles + "MostLinked/"
	EndpointArticlesMostLinkedExpand = EndpointArticles + "MostLinked?expand=1"
	EndpointArticlesNew = EndpointArticles + "New/"
	EndpointArticlesPopular = EndpointArticles + "Popular/"
	EndpointArticlesPopularExpand = EndpointArticles + "Popular?expand=1"
	EndpointArticlesTop = EndpointArticles + "Top/"
	EndpointArticlesTopExpand = EndpointArticles + "Top?expand=1"
	EndpointArticlesTopByHub = EndpointArticles + "TopByHub/" //Honestly, not sure what this is. API says top by views for a hub

	//I will add other things LATER. Library is currently for personal use on a discord bot

	EndpointSearch = EndpointVersion + "Search/"
	EndpointSearchCrossWiki = EndpointSearch + "CrossWiki/"
	EndpointSearchCrossWikiExpand = EndpointSearch + "CrossWiki?expand=1"

	//Not really where this goes but I give no h*cks
	QueryURL = "wiki/Special:Search?query="
)
