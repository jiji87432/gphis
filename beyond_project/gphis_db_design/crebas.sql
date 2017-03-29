/*==============================================================*/
/* DBMS name:      MySQL 5.0                                    */
/* Created on:     2017-03-29 18:07:24                          */
/*==============================================================*/


drop table if exists d_idcard;

drop table if exists t_acceptance;

drop table if exists t_balance;

drop table if exists t_chargetype;

drop table if exists t_detaillist;

drop table if exists t_druginfo;

drop table if exists t_operator;

drop table if exists t_patient;

drop table if exists t_producer;

drop table if exists t_stockinfo;

drop table if exists t_treatment;

/*==============================================================*/
/* Table: d_idcard                                              */
/*==============================================================*/
create table d_idcard
(
   idNum                varchar(6) not null,
   place                varchar(100) not null,
   primary key (idNum)
);

/*==============================================================*/
/* Table: t_acceptance                                          */
/*==============================================================*/
create table t_acceptance
(
   acceptanceID         varchar(14) not null,
   patientID            varchar(14) not null,
   operatorID           varchar(14) not null,
   chiefComplain        varchar(500),
   medicalHistory       varchar(500),
   checks               varchar(500),
   treats               varchar(500),
   diagnose             varchar(500),
   primary key (acceptanceID)
);

/*==============================================================*/
/* Table: t_balance                                             */
/*==============================================================*/
create table t_balance
(
   balanceID            varchar(14) not null,
   typeID               varchar(14) not null,
   totalMount           float not null,
   operatorID           varchar(14) not null,
   acceptanceID         varchar(14),
   primary key (balanceID, typeID)
);

/*==============================================================*/
/* Table: t_chargetype                                          */
/*==============================================================*/
create table t_chargetype
(
   typeID               varchar(14) not null,
   typeName             varchar(30) not null,
   primary key (typeID)
);

/*==============================================================*/
/* Table: t_detaillist                                          */
/*==============================================================*/
create table t_detaillist
(
   acceptanceID         varchar(14) not null,
   itemID               varchar(14) not null,
   itemName             varchar(50) not null,
   unit                 varchar(20) not null,
   unitPrice            float not null,
   mount                float not null,
   totalPrice           float not null,
   operatorID           varchar(14) not null,
   typeID               varchar(14),
   primary key (acceptanceID, itemID)
);

/*==============================================================*/
/* Table: t_druginfo                                            */
/*==============================================================*/
create table t_druginfo
(
   drugID               varchar(14) not null,
   producerID           varchar(14),
   drugName             varchar(50) not null,
   drugMnemonic         varchar(30) not null,
   barcode              varchar(20) not null,
   unit                 varchar(20) not null,
   proSize              varchar(25) not null,
   priceIn              float not null,
   priceOut             float not null,
   primary key (drugID)
);

/*==============================================================*/
/* Table: t_operator                                            */
/*==============================================================*/
create table t_operator
(
   operatorID           varchar(14) not null,
   name                 varchar(20) not null,
   mnemonicCode         varchar(10) not null,
   password             varchar(20) not null,
   primary key (operatorID)
);

/*==============================================================*/
/* Table: t_patient                                             */
/*==============================================================*/
create table t_patient
(
   patientID            varchar(14) not null,
   name                 varchar(20) not null,
   sex                  varchar(6) not null,
   age                  int not null,
   birthday             date,
   idnum                varchar(18),
   address              varchar(100),
   phone                varchar(11),
   primary key (patientID)
);

/*==============================================================*/
/* Table: t_producer                                            */
/*==============================================================*/
create table t_producer
(
   producerID           varchar(14) not null,
   producerName         varchar(100) not null,
   contact              varchar(100),
   primary key (producerID)
);

/*==============================================================*/
/* Table: t_stockinfo                                           */
/*==============================================================*/
create table t_stockinfo
(
   stockID              varchar(14) not null,
   drugID               varchar(14) not null,
   unitPrice            float not null,
   mountIn              float not null,
   stock                float not null,
   operatorID           varchar(14) not null,
   date                 date not null,
   primary key (stockID, drugID)
);

/*==============================================================*/
/* Table: t_treatment                                           */
/*==============================================================*/
create table t_treatment
(
   treatmentID          varchar(14) not null,
   treatmentName        varchar(30) not null,
   unit                 varchar(15) not null,
   unitPrice            float not null,
   primary key (treatmentID)
);

alter table t_acceptance add constraint FK_acceptance_operator foreign key (operatorID)
      references t_operator (operatorID) on delete restrict on update restrict;

alter table t_acceptance add constraint FK_patient_acceptance foreign key (patientID)
      references t_patient (patientID) on delete restrict on update restrict;

alter table t_balance add constraint FK_acc_balance foreign key (acceptanceID)
      references t_acceptance (acceptanceID) on delete restrict on update restrict;

alter table t_balance add constraint FK_balance_chargetype foreign key (typeID)
      references t_chargetype (typeID) on delete restrict on update restrict;

alter table t_balance add constraint FK_balance_operator foreign key (operatorID)
      references t_operator (operatorID) on delete restrict on update restrict;

alter table t_detaillist add constraint FK_detail_acceptance foreign key (acceptanceID)
      references t_acceptance (acceptanceID) on delete restrict on update restrict;

alter table t_detaillist add constraint FK_detail_operator foreign key (operatorID)
      references t_operator (operatorID) on delete restrict on update restrict;

alter table t_detaillist add constraint FK_type_detail foreign key (typeID)
      references t_chargetype (typeID) on delete restrict on update restrict;

alter table t_druginfo add constraint FK_drug_producer foreign key (producerID)
      references t_producer (producerID) on delete restrict on update restrict;

alter table t_stockinfo add constraint FK_stock_drug foreign key (drugID)
      references t_druginfo (drugID) on delete restrict on update restrict;

alter table t_stockinfo add constraint FK_stock_operator foreign key (operatorID)
      references t_operator (operatorID) on delete restrict on update restrict;

