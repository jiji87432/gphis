/*==============================================================*/
/* DBMS name:      MySQL 5.0                                    */
/* Created on:     2017-04-25 15:21:58                          */
/*==============================================================*/


drop table if exists t_bills;

drop table if exists t_diagnose;

drop table if exists t_dispense;

drop table if exists t_dispense_item;

drop table if exists t_drug;

drop table if exists t_drug_item;

drop table if exists t_operator;

drop table if exists t_patient;

/*==============================================================*/
/* Table: t_bills                                               */
/*==============================================================*/
create table t_bills
(
   b_id                 varchar(14) not null,
   diag_id              varchar(14) not null,
   total_expense        float not null,
   o_id                 varchar(14) not null,
   b_date               date not null,
   primary key (b_id)
);

/*==============================================================*/
/* Table: t_diagnose                                            */
/*==============================================================*/
create table t_diagnose
(
   diag_id              varchar(14) not null,
   diag_date            date not null,
   p_id                 varchar(14) not null,
   o_id                 varchar(14) not null,
   drug_item_id         varchar(14),
   dis_item_id          varchar(14),
   diag_complain        varchar(500),
   diag_checks          varchar(500),
   diag_diagnose        varchar(500),
   primary key (diag_id)
);

/*==============================================================*/
/* Table: t_dispense                                            */
/*==============================================================*/
create table t_dispense
(
   dis_id               varchar(14) not null,
   dis_name             varchar(50) not null,
   dis_pinyin           varchar(25) not null,
   dis_unit             varchar(10) not null,
   dis_u_price          float not null,
   primary key (dis_id)
);

/*==============================================================*/
/* Table: t_dispense_item                                       */
/*==============================================================*/
create table t_dispense_item
(
   dis_item_id          varchar(14) not null,
   dis_id               varchar(14) not null,
   mount                float not null,
   primary key (dis_item_id, dis_id)
);

/*==============================================================*/
/* Table: t_drug                                                */
/*==============================================================*/
create table t_drug
(
   drug_id              varchar(14) not null,
   drug_name            varchar(50) not null,
   drug_pinyin          varchar(25) not null,
   drug_barcode         varchar(13) not null,
   drug_unit            varchar(10) not null,
   drug_u_price         float not null,
   drug_spec            varchar(10) not null,
   drug_producer        varchar(50) not null,
   drug_mount           float not null,
   primary key (drug_id)
);

/*==============================================================*/
/* Table: t_drug_item                                           */
/*==============================================================*/
create table t_drug_item
(
   drug_item_id         varchar(14) not null,
   drug_id              varchar(14) not null,
   mount                float not null,
   primary key (drug_item_id, drug_id)
);

/*==============================================================*/
/* Table: t_operator                                            */
/*==============================================================*/
create table t_operator
(
   o_id                 varchar(14) not null,
   o_name               varchar(20) not null,
   o_pinyin             varchar(10) not null,
   o_permit             varchar(100) not null,
   o_pwd                varchar(100) not null,
   primary key (o_id)
);

/*==============================================================*/
/* Table: t_patient                                             */
/*==============================================================*/
create table t_patient
(
   p_id                 varchar(14) not null,
   p_name               varchar(20) not null,
   p_sex                varchar(10) not null,
   p_age                int not null,
   p_birthday           date,
   p_idcard             varchar(18),
   p_address            varchar(150),
   p_contact            varchar(11),
   p_height             int,
   p_weight             int,
   p_gms                varchar(500),
   p_sss                varchar(500),
   p_mxbs               varchar(500),
   primary key (p_id)
);

alter table t_bills add constraint FK_bill_diag foreign key (diag_id)
      references t_diagnose (diag_id) on delete restrict on update restrict;

alter table t_diagnose add constraint FK_diag_dis foreign key (dis_item_id)
      references t_dispense_item (dis_item_id) on delete restrict on update restrict;

alter table t_diagnose add constraint FK_diag_drug_item foreign key (drug_item_id)
      references t_drug_item (drug_item_id) on delete restrict on update restrict;

alter table t_diagnose add constraint FK_diag_op foreign key (o_id)
      references t_operator (o_id) on delete restrict on update restrict;

alter table t_diagnose add constraint FK_diag_patient foreign key (p_id)
      references t_patient (p_id) on delete restrict on update restrict;

alter table t_dispense_item add constraint FK_dis_dis_item foreign key (dis_id)
      references t_dispense (dis_id) on delete restrict on update restrict;

alter table t_drug_item add constraint FK_drug_item_drug foreign key (drug_id)
      references t_drug (drug_id) on delete restrict on update restrict;

