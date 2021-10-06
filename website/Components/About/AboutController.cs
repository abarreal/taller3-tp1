using Microsoft.AspNetCore.Mvc;
using Website.Visits;

namespace Website.About
{
    public class AboutController : WebsiteController
    {
        public AboutController(VisitCounterService visitcount) : base(visitcount)
        {
        }

        [Route("/about")]
        public IActionResult Index()
        {
            return View();
        }

        [Route("/about/legals")]
        public IActionResult Legals()
        {
            return View();
        }

        [Route("/about/offices")]
        public IActionResult Offices()
        {
            return View();
        }
    }
}