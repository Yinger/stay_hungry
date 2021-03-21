using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Configuration;

namespace MicroService.Demo.ServiceInstance.Controllers
{
    [Route("[controller]")]
    [ApiController]
    public class DemoController : ControllerBase
    {
        private IConfiguration _iConfiguration;

        public DemoController(IConfiguration configuration)
        {
            _iConfiguration = configuration;
        }

        [HttpGet]
        public string Get()
        {
            return "Hi i am " + _iConfiguration["port"];
        }
    }
}