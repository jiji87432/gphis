using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace gphis_client_gui
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();
        }

        private void button1_Click(object sender, EventArgs e)
        {
            PatientInsert pi = new PatientInsert();
            pi.Show();
        }

        private void button2_Click(object sender, EventArgs e)
        {
            DoctorWorkspace dw = new DoctorWorkspace();
            dw.Show();
        }

        private void button3_Click(object sender, EventArgs e)
        {
            DiagInsert di = new DiagInsert();
            di.Show();
        }

        private void button4_Click(object sender, EventArgs e)
        {
            DrugInsert di = new DrugInsert();
            di.Show();
        }

        private void button5_Click(object sender, EventArgs e)
        {
            TreatsInsert ti = new TreatsInsert();
            ti.Show();
        }
    }
}
