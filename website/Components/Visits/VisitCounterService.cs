using System.Threading.Tasks;

namespace Website.Visits {
    public interface VisitCounterService
    {
        Task CountVisitAsync(string path);
        Task<int> GetVisitCount();
    }
}