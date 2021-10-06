using System;
using Microsoft.AspNetCore.Mvc;
using Website.Visits;

namespace Website.Home
{
    public class HomeController : WebsiteController
    {

        public HomeController(VisitCounterService visitcount) : base(visitcount)
        {
        }

        [Route("/home")]
        public IActionResult Index()
        {
            var start = DateTime.Now;
            
            // Consume time as if querying a service, a database or something.
            // Let loading jobs take some fixed amount of milliseconds.
            var rand = new Random(Guid.NewGuid().GetHashCode());
            var delay = 100 + rand.Next(0, 100) - 50;

            while (true)
            {
                if ((DateTime.Now - start).TotalMilliseconds > delay)
                {
                    break;
                }
            }

            return View();
        }
    }
}
