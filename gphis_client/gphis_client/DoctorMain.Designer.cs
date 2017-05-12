namespace gphis_client
{
    partial class DoctorMain
    {
        /// <summary>
        /// Required designer variable.
        /// </summary>
        private System.ComponentModel.IContainer components = null;

        /// <summary>
        /// Clean up any resources being used.
        /// </summary>
        /// <param name="disposing">true if managed resources should be disposed; otherwise, false.</param>
        protected override void Dispose(bool disposing)
        {
            if (disposing && (components != null))
            {
                components.Dispose();
            }
            base.Dispose(disposing);
        }

        #region Windows Form Designer generated code

        /// <summary>
        /// Required method for Designer support - do not modify
        /// the contents of this method with the code editor.
        /// </summary>
        private void InitializeComponent()
        {
            this.menuStrip1 = new System.Windows.Forms.MenuStrip();
            this.toolStripMenuItem1 = new System.Windows.Forms.ToolStripMenuItem();
            this.诊疗业务管理ToolStripMenuItem = new System.Windows.Forms.ToolStripMenuItem();
            this.诊所业务ToolStripMenuItem = new System.Windows.Forms.ToolStripMenuItem();
            this.退出ToolStripMenuItem = new System.Windows.Forms.ToolStripMenuItem();
            this.患者信息添加ToolStripMenuItem = new System.Windows.Forms.ToolStripMenuItem();
            this.患者信息维护ToolStripMenuItem = new System.Windows.Forms.ToolStripMenuItem();
            this.menuStrip1.SuspendLayout();
            this.SuspendLayout();
            // 
            // menuStrip1
            // 
            this.menuStrip1.Items.AddRange(new System.Windows.Forms.ToolStripItem[] {
            this.toolStripMenuItem1,
            this.诊疗业务管理ToolStripMenuItem,
            this.诊所业务ToolStripMenuItem,
            this.退出ToolStripMenuItem});
            this.menuStrip1.Location = new System.Drawing.Point(0, 0);
            this.menuStrip1.Name = "menuStrip1";
            this.menuStrip1.Size = new System.Drawing.Size(674, 25);
            this.menuStrip1.TabIndex = 0;
            this.menuStrip1.Text = "menuStrip1";
            // 
            // toolStripMenuItem1
            // 
            this.toolStripMenuItem1.DropDownItems.AddRange(new System.Windows.Forms.ToolStripItem[] {
            this.患者信息添加ToolStripMenuItem,
            this.患者信息维护ToolStripMenuItem});
            this.toolStripMenuItem1.Name = "toolStripMenuItem1";
            this.toolStripMenuItem1.Size = new System.Drawing.Size(68, 21);
            this.toolStripMenuItem1.Text = "患者管理";
            this.toolStripMenuItem1.Click += new System.EventHandler(this.toolStripMenuItem1_Click);
            // 
            // 诊疗业务管理ToolStripMenuItem
            // 
            this.诊疗业务管理ToolStripMenuItem.Name = "诊疗业务管理ToolStripMenuItem";
            this.诊疗业务管理ToolStripMenuItem.Size = new System.Drawing.Size(68, 21);
            this.诊疗业务管理ToolStripMenuItem.Text = "诊疗业务";
            // 
            // 诊所业务ToolStripMenuItem
            // 
            this.诊所业务ToolStripMenuItem.Name = "诊所业务ToolStripMenuItem";
            this.诊所业务ToolStripMenuItem.Size = new System.Drawing.Size(68, 21);
            this.诊所业务ToolStripMenuItem.Text = "诊所业务";
            // 
            // 退出ToolStripMenuItem
            // 
            this.退出ToolStripMenuItem.Name = "退出ToolStripMenuItem";
            this.退出ToolStripMenuItem.Size = new System.Drawing.Size(44, 21);
            this.退出ToolStripMenuItem.Text = "退出";
            this.退出ToolStripMenuItem.Click += new System.EventHandler(this.退出ToolStripMenuItem_Click);
            // 
            // 患者信息添加ToolStripMenuItem
            // 
            this.患者信息添加ToolStripMenuItem.Name = "患者信息添加ToolStripMenuItem";
            this.患者信息添加ToolStripMenuItem.Size = new System.Drawing.Size(152, 22);
            this.患者信息添加ToolStripMenuItem.Text = "患者信息添加";
            this.患者信息添加ToolStripMenuItem.Click += new System.EventHandler(this.患者信息添加ToolStripMenuItem_Click);
            // 
            // 患者信息维护ToolStripMenuItem
            // 
            this.患者信息维护ToolStripMenuItem.Name = "患者信息维护ToolStripMenuItem";
            this.患者信息维护ToolStripMenuItem.Size = new System.Drawing.Size(152, 22);
            this.患者信息维护ToolStripMenuItem.Text = "患者信息维护";
            // 
            // DoctorMain
            // 
            this.AutoScaleDimensions = new System.Drawing.SizeF(6F, 12F);
            this.AutoScaleMode = System.Windows.Forms.AutoScaleMode.Font;
            this.ClientSize = new System.Drawing.Size(674, 426);
            this.Controls.Add(this.menuStrip1);
            this.FormBorderStyle = System.Windows.Forms.FormBorderStyle.FixedDialog;
            this.MainMenuStrip = this.menuStrip1;
            this.Name = "DoctorMain";
            this.StartPosition = System.Windows.Forms.FormStartPosition.CenterScreen;
            this.Text = "医生工作站";
            this.menuStrip1.ResumeLayout(false);
            this.menuStrip1.PerformLayout();
            this.ResumeLayout(false);
            this.PerformLayout();

        }

        #endregion

        private System.Windows.Forms.MenuStrip menuStrip1;
        private System.Windows.Forms.ToolStripMenuItem toolStripMenuItem1;
        private System.Windows.Forms.ToolStripMenuItem 诊疗业务管理ToolStripMenuItem;
        private System.Windows.Forms.ToolStripMenuItem 诊所业务ToolStripMenuItem;
        private System.Windows.Forms.ToolStripMenuItem 退出ToolStripMenuItem;
        private System.Windows.Forms.ToolStripMenuItem 患者信息添加ToolStripMenuItem;
        private System.Windows.Forms.ToolStripMenuItem 患者信息维护ToolStripMenuItem;
    }
}