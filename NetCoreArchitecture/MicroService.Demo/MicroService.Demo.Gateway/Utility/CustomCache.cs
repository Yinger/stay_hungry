﻿using Ocelot.Cache;
using System;
using System.Collections.Generic;
using System.Linq;

namespace MicroService.Demo.Gateway.Utility
{
    /// <summary>
    /// 自定义的Cache 实现IOcelotCache<CachedResponse>
    /// </summary>
    public class CustomCache : IOcelotCache<CachedResponse>
    {
        private class CacheDataModel
        {
            public CachedResponse CachedResponse { get; set; }
            public DateTime Timeout { get; set; }
            public string Region { get; set; }
        }

        private static Dictionary<string, CacheDataModel> CustomCacheDictionary = new
             Dictionary<string, CacheDataModel>();

        /// <summary>
        /// 没做过期清理  所以需要
        /// </summary>
        /// <param name="key"></param>
        /// <param name="value"></param>
        /// <param name="ttl"></param>
        /// <param name="region"></param>
        public void Add(string key, CachedResponse value, TimeSpan ttl, string region)
        {
            Console.WriteLine($"This is {nameof(CustomCache)}.{nameof(Add)}");
            //CustomCacheDictionary.Add(key, new CacheDataModel()
            //{
            //    CachedResponse = value,
            //    Region = region,
            //    Timeout = DateTime.Now.Add(ttl)
            //});
            CustomCacheDictionary[key] = new CacheDataModel()
            {
                CachedResponse = value,
                Region = region,
                Timeout = DateTime.Now.Add(ttl)
            };
        }

        public void AddAndDelete(string key, CachedResponse value, TimeSpan ttl, string region)
        {
            Console.WriteLine($"This is {nameof(CustomCache)}.{nameof(AddAndDelete)}");
            CustomCacheDictionary[key] = new CacheDataModel()
            {
                CachedResponse = value,
                Region = region,
                Timeout = DateTime.Now.Add(ttl)
            };
        }

        public void ClearRegion(string region)
        {
            Console.WriteLine($"This is {nameof(CustomCache)}.{nameof(ClearRegion)}");
            var keyList = CustomCacheDictionary.Where(kv => kv.Value.Region.Equals(region)).Select(kv => kv.Key);
            foreach (var key in keyList)
            {
                CustomCacheDictionary.Remove(key);
            }
        }

        public CachedResponse Get(string key, string region)
        {
            Console.WriteLine($"This is {nameof(CustomCache)}.{nameof(Get)}");
            if (CustomCacheDictionary.ContainsKey(key) && CustomCacheDictionary[key] != null
                && CustomCacheDictionary[key].Timeout > DateTime.Now
                && CustomCacheDictionary[key].Region.Equals(region))
            {
                return CustomCacheDictionary[key].CachedResponse;
            }
            else
                return null;
        }
    }
}