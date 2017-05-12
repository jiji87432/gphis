using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace gphis_client
{
    public partial class PatientInsert : Form
    {
        public PatientInsert()
        {
            InitializeComponent();
        }

        private void PatientInsert_Load(object sender, EventArgs e)
        {
            button2_Click(null, null);
        }

        public string GetID()
        {
            string year = DateTime.Now.Year.ToString().Substring(2, 2);
            string month = DateTime.Now.Month < 10 ? "0" + DateTime.Now.Month.ToString() : DateTime.Now.Month.ToString();
            string day = DateTime.Now.Day < 10 ? "0" + DateTime.Now.Day.ToString() : DateTime.Now.Day.ToString();
            string hour = DateTime.Now.Hour < 10 ? "0" + DateTime.Now.Hour.ToString() : DateTime.Now.Hour.ToString();
            string minute = DateTime.Now.Minute < 10 ? "0" + DateTime.Now.Minute.ToString() : DateTime.Now.Minute.ToString();
            string second = DateTime.Now.Second < 10 ? "0" + DateTime.Now.Second.ToString() : DateTime.Now.Second.ToString();
            Random r = new Random();
            int rand = r.Next(0, 99);
            string random = rand < 10 ? "0" + rand.ToString() : rand.ToString();
            return string.Format("{0}{1}{2}{3}{4}{5}{6}", year, month, day, hour, minute, second, random);
        }

        private void button2_Click(object sender, EventArgs e)
        {
            NameTextBox.Text = "";
            AgeTextBox.Text = "";
            IDCardTextBox.Text = "";
            SexComboBox.SelectedIndex = 0;
            BirthdayBox.Value = DateTime.Now;
            AddressBox.Text = "";
            HeightBox.Text = "";
            WeightBox.Text = "";
            ContactBox.Text = "";
            GMSBox.Text = "";
            SSSBox.Text = "";
            MXBSBox.Text = "";
            PatientIDLabel.Text = GetID();
        }

        private void button1_Click(object sender, EventArgs e)
        {

        }
    }
}
