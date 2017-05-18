using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Newtonsoft.Json;
using System.IO;

namespace HTTPTest
{
    public class JSONHelper
    {
        public static string Serialize(object o)
        {
            var json = JsonConvert.SerializeObject(o);
            return json;
        }

        public static T DeserializeToObject<T>(string json) where T : class
        {
            JsonSerializer s = new JsonSerializer();
            StringReader r = new StringReader(json);
            object o = s.Deserialize(new JsonTextReader(r), typeof(T));
            T t = o as T;
            return t;
        }
    }
}
