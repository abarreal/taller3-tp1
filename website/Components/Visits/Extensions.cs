using System.Net.Http;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;

namespace Website.Visits
{
	public static class Extensions
	{
		public static IServiceCollection AddPubSubVisitCounterService(
			this IServiceCollection services,
			IConfiguration config,
			HttpClient client)
		{
			// Read visit counter service configuration.
			var visitcounterConfig = new VisitCounterConfig();
            var visitcounterPubSub = new PubSubVisitCounterConfig();
			// Read the VisitCounterService section from JSON configuration and use that
			// to initialize the PubSubVisitCounterConfig object.
			config.GetSection("VisitCounterService").Bind(visitcounterConfig);
            config.GetSection("PubSubVisitCounterService").Bind(visitcounterPubSub);
            // Register the visit counter service.
            return services.AddSingleton<VisitCounterService>(x =>
                new PubSubVisitCounterService(client, visitcounterConfig, visitcounterPubSub));
		}

	}
}