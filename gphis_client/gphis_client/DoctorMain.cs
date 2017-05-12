using gphis_client.Tables;
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
    public partial class DoctorMain : Form
    {
        public DoctorMain(T_Operator o)
        {
            InitializeComponent();
            op = o;
        }

        private T_Operator op;

        private void toolStripMenuItem1_Click(object sender, EventArgs e)
        {

        }

        private void 退出ToolStripMenuItem_Click(object sender, EventArgs e)
        {
            this.Close();
        }

        private void 患者信息添加ToolStripMenuItem_Click(object sender, EventArgs e)
        {
            PatientInsert pi = new PatientInsert();
            pi.Show();
        }
    }
}
