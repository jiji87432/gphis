/*==============================================================*/
/* DBMS name:      MySQL 5.0                                    */
/* Created on:     2017/5/3 21:19:17                            */
/*==============================================================*/


drop table if exists bills;

drop table if exists diagnose;

drop table if exists dispense;

drop table if exists dispense_item;

drop table if exists drug;

drop table if exists drug_item;

drop table if exists operator;

drop table if exists patient;

/*==============================================================*/
/* Table: bills                                                 */
/*==============================================================*/
create table bills
(
   b_id                 varchar(14) not null,
   diag_id              varchar(14) not null,
   total_expense        float not null,
   o_id                 varchar(14) not null,
   b_date               datetime not null,
   primary key (b_id)
);

/*==============================================================*/
/* Table: diagnose                                              */
/*==============================================================*/
create table diagnose
(
   diag_id              varchar(14) not null,
   diag_date            datetime not null,
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
/* Table: dispense                                              */
/*==============================================================*/
create table dispense
(
   dis_id               varchar(14) not null,
   dis_name             varchar(50) not null,
   dis_pinyin           varchar(25) not null,
   dis_unit             varchar(10) not null,
   dis_u_price          float not null,
   primary key (dis_id)
);

/*==============================================================*/
/* Table: dispense_item                                         */
/*==============================================================*/
create table dispense_item
(
   dis_item_id          varchar(14) not null,
   dis_id               varchar(14) not null,
   mount                float not null,
   primary key (dis_item_id, dis_id)
);

/*==============================================================*/
/* Table: drug                                                  */
/*==============================================================*/
create table drug
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
/* Table: drug_item                                             */
/*==============================================================*/
create table drug_item
(
   drug_item_id         varchar(14) not null,
   drug_id              varchar(14) not null,
   mount                float not null,
   primary key (drug_item_id, drug_id)
);

/*==============================================================*/
/* Table: operator                                              */
/*==============================================================*/
create table operator
(
   o_id                 varchar(14) not null,
   o_name               varchar(20) not null,
   o_pinyin             varchar(10) not null,
   o_permit             varchar(100) not null,
   o_pwd                varchar(100) not null,
   primary key (o_id)
);

/*==============================================================*/
/* Table: patient                                               */
/*==============================================================*/
create table patient
(
   p_id                 varchar(14) not null,
   p_name               varchar(20) not null,
   p_sex                varchar(10) not null,
   p_age                int not null,
   p_birthday           date,
   p_idcard             varchar(18),
   p_address            varchar(150),
   p_contact            varchar(11),
   p_height             float,
   p_weight             float,
   p_gms                varchar(500),
   p_sss                varchar(500),
   p_mxbs               varchar(500),
   primary key (p_id)
);

alter table bills add constraint FK_bill_diag foreign key (diag_id)
      references diagnose (diag_id) on delete restrict on update restrict;

alter table diagnose add constraint FK_diag_dis foreign key (dis_item_id)
      references dispense_item (dis_item_id) on delete restrict on update restrict;

alter table diagnose add constraint FK_diag_drug_item foreign key (drug_item_id)
      references drug_item (drug_item_id) on delete restrict on update restrict;

alter table diagnose add constraint FK_diag_op foreign key (o_id)
      references operator (o_id) on delete restrict on update restrict;

alter table diagnose add constraint FK_diag_patient foreign key (p_id)
      references patient (p_id) on delete restrict on update restrict;

alter table dispense_item add constraint FK_dis_dis_item foreign key (dis_id)
      references dispense (dis_id) on delete restrict on update restrict;

alter table drug_item add constraint FK_drug_item_drug foreign key (drug_id)
      references drug (drug_id) on delete restrict on update restrict;

