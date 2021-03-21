using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;
using MicroService.Demo.Client.Models;
using System.Net.Http;
using Consul;

namespace MicroService.Demo.Client.Controllers
{
    public class HomeController : Controller
    {
        private readonly ILogger<HomeController> _logger;
        private static long iTotalCount = 0;

        public HomeController(ILogger<HomeController> logger)
        {
            _logger = logger;
        }

        public IActionResult Index()
        {
            //string url = "http://DemoService/demo";

            //ConsulClient client = new ConsulClient(c =>
            //{
            //    c.Address = new Uri("http://8.131.87.88:8083/");//consul地址
            //    c.Datacenter = "cc-1";
            //});

            //var response = client.Agent.Services().Result.Response;
            //foreach (var item in response)
            //{
            //    Console.WriteLine("***************************************");
            //    Console.WriteLine(item.Key);
            //    var service = item.Value;
            //    Console.WriteLine($"{service.Address}--{service.Port}--{service.Service}");
            //    Console.WriteLine("***************************************");
            //}

            //Uri uri = new Uri(url);
            //string groupName = uri.Host;
            //AgentService agentService = null;

            //var serviceDictionary = response.Where(s => s.Value.Service.Equals(groupName, StringComparison.OrdinalIgnoreCase)).ToArray();//获取的全部服务实例信息

            //////轮询策略
            ////if (iTotalCount > 10000)
            ////    iTotalCount = 0;
            ////agentService = serviceDictionary[iTotalCount++ % serviceDictionary.Length].Value;

            ////随机策略
            //agentService = serviceDictionary[new Random().Next(0, 1000) % serviceDictionary.Length].Value;

            ////权重策略--不同的服务器处理能力不同，按能力分配
            ////serviceDictionary[0].Value.Tags//获取权重1
            ////https://blog.csdn.net/qq_26900081/article/details/108052868

            //url = $"{uri.Scheme}://{agentService.Address}:{agentService.Port}{uri.PathAndQuery}";

            //string content = InvokeApi(url);
            //base.ViewBag.MSG = content;
            //Console.WriteLine($"This is {url} Invoke");

            string url = "http://8.131.87.88:6298/T/demo";
            string content = InvokeApi(url);
            ViewBag.MSG = content;
            Console.WriteLine($"This is {url} Invoke");

            return View();
        }

        public IActionResult Privacy()
        {
            return View();
        }

        [ResponseCache(Duration = 0, Location = ResponseCacheLocation.None, NoStore = true)]
        public IActionResult Error()
        {
            return View(new ErrorViewModel { RequestId = Activity.Current?.Id ?? HttpContext.TraceIdentifier });
        }

        public static string InvokeApi(string url)
        {
            using (HttpClient httpClient = new HttpClient())
            {
                HttpRequestMessage message = new HttpRequestMessage();
                message.Method = HttpMethod.Get;
                message.RequestUri = new Uri(url);
                var result = httpClient.SendAsync(message).Result;
                string content = result.Content.ReadAsStringAsync().Result;
                return content;
            }
        }
    }
}