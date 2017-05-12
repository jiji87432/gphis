using gphis_client.Tables;
using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Security.Cryptography;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace gphis_client
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();
        }

        private T_Operator op;
        private string IP;
        private string Port;

        private void button2_Click(object sender, EventArgs e)
        {
            this.Close();
        }

        private void button1_Click(object sender, EventArgs e)
        {
            var conn = ConnectionTest(IPTextBox.Text, PortTextBox.Text);
            if (conn == true)
            {
                string pwd = StrToMD5(this.PasswordTextBox.Text);
                var o = GetOne(this.UserNameTextBox.Text);
                if (o != null)
                {
                    if (o.o_pwd == pwd)
                    {
                        DoctorMain dm = new DoctorMain(o);
                        dm.Show();
                    }
                }
            }
        }

        public static string StrToMD5(string str)
        {
            byte[] data = Encoding.UTF8.GetBytes(str);
            MD5 md5 = new MD5CryptoServiceProvider();
            byte[] OutBytes = md5.ComputeHash(data);

            string OutString = "";
            for (int i = 0; i < OutBytes.Length; i++)
            {
                OutString += OutBytes[i].ToString("x2");
            }
            // return OutString.ToUpper();
            return OutString.ToLower();
        }

        private T_Operator GetOne(string pinyin)
        {
            T_Operator o = new T_Operator();
            o.o_pinyin = UserNameTextBox.Text;
            var data = JsonHelper.Serialize(o);
            var hresp = HttpHelper.PostJSON(string.Format("http://{0}:{1}/operator/get_one", IP, Port), data);
            var resp = HttpHelper.GetResponseString(hresp);
            var r = JsonHelper.DeserializeToObject<Response>(resp);
            if (r.result == true)
            {
                return JsonHelper.DeserializeToObject<T_Operator>(JsonHelper.Serialize(r.reason));
            }
            else
            {
                return null;
            }
        }

        private bool ConnectionTest(string ip, string port)
        {
            bool ret = false;
            try
            {
                var resp = HttpHelper.CreateGetHttpResponse("http://" + ip + ":" + port + "/", 5, "", null);
                var str = HttpHelper.GetResponseString(resp);
                var response = JsonHelper.DeserializeToObject<Response>(str);
                if (response.result == true)
                {
                    ret = true;
                    IP = ip;
                    Port = port;
                }
            }
            catch(Exception e)
            {
                MessageBox.Show(e.Message);
            }
            return ret;
        }
    }
}
