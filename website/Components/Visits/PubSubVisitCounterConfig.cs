namespace Website.Visits
{
    public class PubSubVisitCounterConfig
    {
        public bool IsPublishingEnabled { get; set; }
        public string ProjectId { get; set; }
        public string TopicId { get; set; }
    }
}