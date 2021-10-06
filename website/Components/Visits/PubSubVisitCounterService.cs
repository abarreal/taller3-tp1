using System.Net.Http;
using System.Threading.Tasks;
using Google.Apis.Logging;
using Google.Cloud.PubSub.V1;
using Microsoft.Extensions.Logging;
using Newtonsoft.Json;

namespace Website.Visits
{
    public class PubSubVisitCounterService : VisitCounterService
    {
        private readonly HttpClient _http;
        private readonly VisitCounterConfig _config;
        private readonly PubSubVisitCounterConfig _pubsubConfig;
        private readonly PublisherClient _publisher;

        public PubSubVisitCounterService(
            HttpClient http,
            VisitCounterConfig config,
            PubSubVisitCounterConfig pubsubConfig)
        {
            this._http = http;
            this._config = config;
            this._pubsubConfig = pubsubConfig;
            // Do nothing if publishing is not enabled.
            if (!pubsubConfig.IsPublishingEnabled) {
                return;
            }
            System.Console.WriteLine($"Visit Counter Project: {pubsubConfig.ProjectId}");
            System.Console.WriteLine($"Visit Counter Topic: {pubsubConfig.TopicId}");
            // Create topic reference.
            var topic = TopicName.FromProjectTopic(
                pubsubConfig.ProjectId,
                pubsubConfig.TopicId);
            // Instantiate publisher that pushes messages to the topic.
            this._publisher = PublisherClient.Create(topic);
        }

        public async Task CountVisitAsync(string path)
        {
            // Do nothing if publishing is not enabled.
            if (!this._pubsubConfig.IsPublishingEnabled) {
                return;
            }
            // Generate the JSON message. The message is currently very simple,
            // so we generate it from a template string.
            var message = $"{{ \"path\": \"{path}\" }}";
            System.Console.WriteLine($"Pushing visit message {message}");
            await this._publisher.PublishAsync(message);
            System.Console.WriteLine("Visit message published");
        }

        public async Task<int> GetVisitCount() {
            // Construct the path to the API endpoint.
            var path = $"{this._config.VisitCounterApiBase}{this._config.GetCountEndpoint}";
            // Execute API call and get the count.
            System.Console.WriteLine($"Retrieving visit count from {path}");
            var response = (await this._http.GetAsync(path)).EnsureSuccessStatusCode();
            // Read the body of the response.
            string body = await response.Content.ReadAsStringAsync();
            System.Console.WriteLine($"Received {body}");
            // Parse the JSON response and return the count.
            return JsonConvert.DeserializeObject<VisitCountResponse>(body).Count;
        }

        private class VisitCountResponse {
            public int Count { get; set; }
        }

    }
}