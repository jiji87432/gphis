using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace HTTPTest
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();
        }

        private void button1_Click(object sender, EventArgs e)
        {
            //var str = HTTPHelper.GetResponseString(HTTPHelper.CreateGetHttpResponse("http://192.168.142.128:10160/drug/all", 10, "", null));
            //var o = JSONHelper.DeserializeToObject<Response>(str);
            //richTextBox1.Text = o.result.ToString();
            T_Operator o = new T_Operator();
            o.o_pinyin = "cf";
            var str = JSONHelper.Serialize(o);
            var resp = HTTPHelper.GetResponseString(HTTPHelper.PostJSON("http://192.168.142.128:10160/operator/get_one", str));
            var r = JSONHelper.DeserializeToObject<Response>(resp);
            var hh = JSONHelper.DeserializeToObject<T_Operator>(JSONHelper.Serialize(r.reason));
            richTextBox1.Text = hh.o_pinyin;
        }
    }
}
