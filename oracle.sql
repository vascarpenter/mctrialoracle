--------------------------------------------------------
--  ファイルを作成しました - 月曜日-6月-29-2020   
--------------------------------------------------------
--------------------------------------------------------
--  DDL for Table EVENTS
--------------------------------------------------------

  CREATE TABLE "ADMIN"."EVENTS" 
   (	"ID" NUMBER GENERATED BY DEFAULT ON NULL AS IDENTITY MINVALUE 1 MAXVALUE 9999999999999999999999999999 INCREMENT BY 1 START WITH 1 CACHE 20 NOORDER  NOCYCLE  NOKEEP  NOSCALE, 
	"HOSPITAL_ID" NUMBER, 
	"SERIALID" NUMBER, 
	"EVENTID" NUMBER, 
	"DATE" DATE DEFAULT NULL, 
	"ALIVE" NUMBER DEFAULT NULL, 
	"DROPOUT" NUMBER DEFAULT NULL, 
	"MACCE" NUMBER DEFAULT NULL, 
	"BH" NUMBER DEFAULT NULL, 
	"BW" NUMBER DEFAULT NULL, 
	"SBP" NUMBER DEFAULT NULL, 
	"DBP" NUMBER DEFAULT NULL, 
	"HR" NUMBER DEFAULT NULL, 
	"EVENT" NVARCHAR2(80) COLLATE "USING_NLS_COMP" DEFAULT NULL, 
	"DIFFDAYS" NUMBER DEFAULT NULL
   )  DEFAULT COLLATION "USING_NLS_COMP" SEGMENT CREATION IMMEDIATE 
  PCTFREE 10 PCTUSED 40 INITRANS 1 MAXTRANS 255 
 NOCOMPRESS LOGGING
  STORAGE(INITIAL 65536 NEXT 1048576 MINEXTENTS 1 MAXEXTENTS 2147483645
  PCTINCREASE 0 FREELISTS 1 FREELIST GROUPS 1
  BUFFER_POOL DEFAULT FLASH_CACHE DEFAULT CELL_FLASH_CACHE DEFAULT)
  TABLESPACE "DATA" ;
--------------------------------------------------------
--  DDL for Table HOSPITALS
--------------------------------------------------------

  CREATE TABLE "ADMIN"."HOSPITALS" 
   (	"ID" NUMBER GENERATED BY DEFAULT ON NULL AS IDENTITY MINVALUE 1 MAXVALUE 9999999999999999999999999999 INCREMENT BY 1 START WITH 1 CACHE 20 NOORDER  NOCYCLE  NOKEEP  NOSCALE , 
	"HOSPITAL_ID" NUMBER, 
	"NAME" NVARCHAR2(80) COLLATE "USING_NLS_COMP", 
	"CREATED_AT" DATE, 
	"USERID" NVARCHAR2(20) COLLATE "USING_NLS_COMP", 
	"USERPASS" NVARCHAR2(80) COLLATE "USING_NLS_COMP", 
	"MAILADDRESS" NVARCHAR2(80) COLLATE "USING_NLS_COMP" DEFAULT NULL
   )  DEFAULT COLLATION "USING_NLS_COMP" SEGMENT CREATION IMMEDIATE 
  PCTFREE 10 PCTUSED 40 INITRANS 1 MAXTRANS 255 
 NOCOMPRESS LOGGING
  STORAGE(INITIAL 65536 NEXT 1048576 MINEXTENTS 1 MAXEXTENTS 2147483645
  PCTINCREASE 0 FREELISTS 1 FREELIST GROUPS 1
  BUFFER_POOL DEFAULT FLASH_CACHE DEFAULT CELL_FLASH_CACHE DEFAULT)
  TABLESPACE "DATA" ;
--------------------------------------------------------
--  DDL for Table PATIENTS
--------------------------------------------------------

  CREATE TABLE "ADMIN"."PATIENTS" 
   (	"ID" NUMBER GENERATED BY DEFAULT ON NULL AS IDENTITY MINVALUE 1 MAXVALUE 9999999999999999999999999999 INCREMENT BY 1 START WITH 1 CACHE 20 NOORDER  NOCYCLE  NOKEEP  NOSCALE , 
	"PATIENT_ID" NVARCHAR2(40) COLLATE "USING_NLS_COMP", 
	"HOSPITAL_ID" NUMBER, 
	"SERIALID" NUMBER, 
	"TRIALGROUP" NUMBER, 
	"INITIAL" NVARCHAR2(40) COLLATE "USING_NLS_COMP" DEFAULT NULL, 
	"BIRTHDATE" DATE DEFAULT NULL, 
	"FEMALE" NUMBER DEFAULT NULL, 
	"AGE" NUMBER DEFAULT NULL, 
	"ALLOWDATE" DATE DEFAULT NULL, 
	"STARTDATE" DATE DEFAULT NULL, 
	"DROPDATE" DATE DEFAULT NULL, 
	"MACCEDATE" DATE DEFAULT NULL, 
	"DEADDATE" DATE DEFAULT NULL, 
	"FINISHDATE" DATE DEFAULT NULL, 
	"DIFFDAYS" NUMBER DEFAULT NULL
   )  DEFAULT COLLATION "USING_NLS_COMP" SEGMENT CREATION IMMEDIATE 
  PCTFREE 10 PCTUSED 40 INITRANS 1 MAXTRANS 255 
 NOCOMPRESS LOGGING
  STORAGE(INITIAL 65536 NEXT 1048576 MINEXTENTS 1 MAXEXTENTS 2147483645
  PCTINCREASE 0 FREELISTS 1 FREELIST GROUPS 1
  BUFFER_POOL DEFAULT FLASH_CACHE DEFAULT CELL_FLASH_CACHE DEFAULT)
  TABLESPACE "DATA" ;
REM INSERTING into ADMIN.EVENTS
SET DEFINE OFF;
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (1,7,1,1,to_date('19-03-04','RR-MM-DD'),1,0,0,170,50,120,80,60,'test1',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (2,7,2,1,to_date('19-03-04','RR-MM-DD'),1,0,0,160,80,130,70,70,'test1',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (3,7,2,2,to_date('19-03-05','RR-MM-DD'),1,0,0,150,80,130,70,80,'test2',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (4,7,1,2,to_date('19-03-05','RR-MM-DD'),1,0,0,150,80,130,80,70,'test2',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (5,7,1,3,to_date('19-03-06','RR-MM-DD'),1,0,0,170,65,130,70,60,'３番目',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (7,8,3,1,to_date('19-03-04','RR-MM-DD'),1,0,0,153,67,130,80,60,'event',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (77,7,1,4,to_date('19-03-07','RR-MM-DD'),1,0,0,170,60,140,70,60,'event-7-1-3',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (78,7,1,5,to_date('19-03-13','RR-MM-DD'),1,0,0,170,60,140,70,60,'event-7-1-3',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (79,7,1,6,to_date('19-03-23','RR-MM-DD'),1,0,0,170,60,140,70,60,'event-7-1-3',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (80,7,1,7,to_date('19-04-02','RR-MM-DD'),1,0,0,170,60,140,70,60,'event-7-1-3',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (81,7,1,8,to_date('19-04-12','RR-MM-DD'),1,0,0,170,60,140,70,60,'event-7-1-3',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (82,7,1,9,to_date('19-04-22','RR-MM-DD'),1,0,0,170,60,140,70,60,'event-7-1-3',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (83,7,1,10,to_date('19-05-02','RR-MM-DD'),1,0,0,170,60,140,70,60,'event-7-1-3',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (84,7,1,11,to_date('19-05-12','RR-MM-DD'),1,0,0,170,60,140,70,60,'event-7-1-3',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (85,7,1,12,to_date('19-05-22','RR-MM-DD'),1,0,1,170,60,140,70,60,'event-7-1-3',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (86,7,2,3,to_date('19-03-07','RR-MM-DD'),1,0,0,170,60,140,70,60,'event-7-2-2',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (87,7,2,4,to_date('19-03-13','RR-MM-DD'),1,0,0,170,60,140,70,60,'event-7-2-2',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (88,7,2,5,to_date('19-03-23','RR-MM-DD'),1,0,0,170,60,140,70,60,'event-7-2-2',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (89,7,2,6,to_date('19-04-02','RR-MM-DD'),1,0,0,170,60,140,70,60,'event-7-2-2',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (90,7,2,7,to_date('19-04-12','RR-MM-DD'),1,0,0,170,60,140,70,60,'event-7-2-2',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (91,7,2,8,to_date('19-04-22','RR-MM-DD'),1,0,0,170,60,140,70,60,'event-7-2-2',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (92,7,2,9,to_date('19-05-02','RR-MM-DD'),1,0,0,170,60,140,70,60,'event-7-2-2',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (93,7,2,10,to_date('19-05-12','RR-MM-DD'),1,0,0,170,60,140,70,60,'event-7-2-2',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (94,7,2,11,to_date('19-05-22','RR-MM-DD'),1,0,0,170,60,140,70,60,'event-7-2-2',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (95,7,3,1,to_date('19-03-03','RR-MM-DD'),1,0,0,170,60,140,70,60,'event-7-3-0',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (96,7,3,2,to_date('19-03-13','RR-MM-DD'),1,0,0,170,60,140,70,60,'event-7-3-0',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (97,7,3,3,to_date('19-03-23','RR-MM-DD'),1,0,0,170,60,140,70,60,'event-7-3-0',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (98,7,3,4,to_date('19-04-02','RR-MM-DD'),1,0,0,170,60,140,70,60,'event-7-3-0',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (99,7,3,5,to_date('19-04-12','RR-MM-DD'),1,0,0,170,60,140,70,60,'event-7-3-0',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (100,7,3,6,to_date('19-04-22','RR-MM-DD'),1,0,0,170,60,140,70,60,'event-7-3-0',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (101,7,3,7,to_date('19-05-02','RR-MM-DD'),1,0,0,170,60,140,70,60,'event-7-3-0',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (102,7,3,8,to_date('19-05-12','RR-MM-DD'),1,0,0,170,60,140,70,60,'event-7-3-0',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (103,7,3,9,to_date('19-05-22','RR-MM-DD'),1,0,0,170,60,140,70,60,'event-7-3-0',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (104,8,3,2,to_date('19-03-03','RR-MM-DD'),1,0,0,170,60,140,70,60,'event-8-3-1',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (105,8,3,3,to_date('19-03-13','RR-MM-DD'),1,0,0,170,60,140,70,60,'event-8-3-1',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (106,8,3,4,to_date('19-03-23','RR-MM-DD'),1,0,0,170,60,140,70,60,'event-8-3-1',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (107,8,3,5,to_date('19-04-02','RR-MM-DD'),1,0,0,170,60,140,70,60,'event-8-3-1',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (108,8,3,6,to_date('19-04-12','RR-MM-DD'),1,0,0,170,60,140,70,60,'event-8-3-1',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (109,8,3,7,to_date('19-04-22','RR-MM-DD'),1,0,0,170,60,140,70,60,'event-8-3-1',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (110,8,3,8,to_date('19-05-02','RR-MM-DD'),0,0,0,170,60,140,70,60,'event-8-3-1',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (111,7,4,1,to_date('19-03-03','RR-MM-DD'),1,0,0,170,60,140,70,60,'event-7-4-0',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (112,7,4,2,to_date('19-03-13','RR-MM-DD'),1,0,0,170,60,140,70,60,'event-7-4-0',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (113,7,4,3,to_date('19-03-23','RR-MM-DD'),1,0,0,170,60,140,70,60,'event-7-4-0',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (114,7,4,4,to_date('19-04-02','RR-MM-DD'),1,0,0,170,60,140,70,60,'event-7-4-0',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (115,7,4,5,to_date('19-04-12','RR-MM-DD'),1,0,1,170,60,140,70,60,'event-7-4-0',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (116,7,5,1,to_date('19-03-03','RR-MM-DD'),1,0,0,170,60,140,70,60,'event-7-5-0',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (117,7,5,2,to_date('19-03-13','RR-MM-DD'),1,0,0,170,60,140,70,60,'event-7-5-0',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (118,7,5,3,to_date('19-03-23','RR-MM-DD'),1,0,0,170,60,140,70,60,'event-7-5-0',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (119,7,5,4,to_date('19-04-02','RR-MM-DD'),1,0,1,170,60,140,70,60,'event-7-5-0',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (120,7,6,1,to_date('19-03-03','RR-MM-DD'),1,0,0,170,60,140,70,60,'event-7-6-0',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (121,7,6,2,to_date('19-03-13','RR-MM-DD'),1,0,0,170,60,140,70,60,'event-7-6-0',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (122,7,6,3,to_date('20-03-03','RR-MM-DD'),0,0,0,170,60,140,70,60,'event-7-6-0',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (123,8,2,1,to_date('19-03-03','RR-MM-DD'),1,0,0,170,60,140,70,60,'event-8-2-0',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (124,8,2,2,to_date('19-03-13','RR-MM-DD'),1,0,0,170,60,140,70,60,'event-8-2-0',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (125,8,2,3,to_date('19-03-23','RR-MM-DD'),1,0,0,170,60,140,70,60,'event-8-2-0',null);
Insert into ADMIN.EVENTS (ID,HOSPITAL_ID,SERIALID,EVENTID,"DATE",ALIVE,DROPOUT,MACCE,BH,BW,SBP,DBP,HR,EVENT,DIFFDAYS) values (126,8,2,4,to_date('19-04-02','RR-MM-DD'),1,0,0,170,60,140,70,60,'event-8-2-0',null);
REM INSERTING into ADMIN.HOSPITALS
SET DEFINE OFF;
Insert into ADMIN.HOSPITALS (ID,HOSPITAL_ID,NAME,CREATED_AT,USERID,USERPASS,MAILADDRESS) values (1,7,'サンプル病院(pass test)',to_date('20-01-30','RR-MM-DD'),'test','$2a$10$HuKg5hNOrQvlNn1p6Ck00O45fk1MoL4gsikUbbnvCz7bsAGWEUa.y','test@gmail.com');
Insert into ADMIN.HOSPITALS (ID,HOSPITAL_ID,NAME,CREATED_AT,USERID,USERPASS,MAILADDRESS) values (2,8,'サンプル病院２(pass test2)',to_date('20-02-02','RR-MM-DD'),'test2','$2b$10$lKqbAlp39v1LNPR2tpHzMOpCbatpxlOdEqqr2QhsBdid2ATEAk0zu','test2@me.com');
Insert into ADMIN.HOSPITALS (ID,HOSPITAL_ID,NAME,CREATED_AT,USERID,USERPASS,MAILADDRESS) values (3,1,'admin',to_date('20-01-30','RR-MM-DD'),'admin','$2a$10$1BzsC7cdT.pvQ4fgyIWL..dQP3H2X0/ynPKhL0sbl5kEcPzqH4Ffy','admintest@gmail.com');
REM INSERTING into ADMIN.PATIENTS
SET DEFINE OFF;
Insert into ADMIN.PATIENTS (ID,PATIENT_ID,HOSPITAL_ID,SERIALID,TRIALGROUP,"INITIAL",BIRTHDATE,FEMALE,AGE,ALLOWDATE,STARTDATE,DROPDATE,MACCEDATE,DEADDATE,FINISHDATE,DIFFDAYS) values (1,'1',7,1,0,'A.I.',to_date('70-03-04','RR-MM-DD'),0,50,to_date('19-03-03','RR-MM-DD'),to_date('19-03-04','RR-MM-DD'),to_date('19-05-22','RR-MM-DD'),to_date('19-05-22','RR-MM-DD'),null,null,null);
Insert into ADMIN.PATIENTS (ID,PATIENT_ID,HOSPITAL_ID,SERIALID,TRIALGROUP,"INITIAL",BIRTHDATE,FEMALE,AGE,ALLOWDATE,STARTDATE,DROPDATE,MACCEDATE,DEADDATE,FINISHDATE,DIFFDAYS) values (2,'2',7,2,0,'B.C',to_date('50-05-06','RR-MM-DD'),0,69,to_date('19-03-03','RR-MM-DD'),to_date('19-03-03','RR-MM-DD'),null,null,null,to_date('20-06-27','RR-MM-DD'),null);
Insert into ADMIN.PATIENTS (ID,PATIENT_ID,HOSPITAL_ID,SERIALID,TRIALGROUP,"INITIAL",BIRTHDATE,FEMALE,AGE,ALLOWDATE,STARTDATE,DROPDATE,MACCEDATE,DEADDATE,FINISHDATE,DIFFDAYS) values (3,'3',7,3,0,'D.E',to_date('60-01-23','RR-MM-DD'),0,50,to_date('19-03-03','RR-MM-DD'),to_date('19-03-03','RR-MM-DD'),null,null,null,to_date('20-06-27','RR-MM-DD'),null);
Insert into ADMIN.PATIENTS (ID,PATIENT_ID,HOSPITAL_ID,SERIALID,TRIALGROUP,"INITIAL",BIRTHDATE,FEMALE,AGE,ALLOWDATE,STARTDATE,DROPDATE,MACCEDATE,DEADDATE,FINISHDATE,DIFFDAYS) values (4,'S2-001',8,3,1,'E.C',to_date('77-04-05','RR-MM-DD'),1,42,to_date('19-03-03','RR-MM-DD'),to_date('19-03-03','RR-MM-DD'),null,null,to_date('19-05-02','RR-MM-DD'),null,null);
Insert into ADMIN.PATIENTS (ID,PATIENT_ID,HOSPITAL_ID,SERIALID,TRIALGROUP,"INITIAL",BIRTHDATE,FEMALE,AGE,ALLOWDATE,STARTDATE,DROPDATE,MACCEDATE,DEADDATE,FINISHDATE,DIFFDAYS) values (5,'7',7,4,0,'T.T',to_date('70-01-01','RR-MM-DD'),0,50,to_date('19-03-03','RR-MM-DD'),to_date('19-03-03','RR-MM-DD'),to_date('19-04-12','RR-MM-DD'),to_date('19-04-12','RR-MM-DD'),null,null,null);
Insert into ADMIN.PATIENTS (ID,PATIENT_ID,HOSPITAL_ID,SERIALID,TRIALGROUP,"INITIAL",BIRTHDATE,FEMALE,AGE,ALLOWDATE,STARTDATE,DROPDATE,MACCEDATE,DEADDATE,FINISHDATE,DIFFDAYS) values (6,'8',7,5,0,'F.G',to_date('20-01-01','RR-MM-DD'),0,100,to_date('19-03-03','RR-MM-DD'),to_date('19-03-03','RR-MM-DD'),to_date('19-04-02','RR-MM-DD'),to_date('19-04-02','RR-MM-DD'),null,null,null);
Insert into ADMIN.PATIENTS (ID,PATIENT_ID,HOSPITAL_ID,SERIALID,TRIALGROUP,"INITIAL",BIRTHDATE,FEMALE,AGE,ALLOWDATE,STARTDATE,DROPDATE,MACCEDATE,DEADDATE,FINISHDATE,DIFFDAYS) values (7,'6-6',7,6,0,'I.D',to_date('70-01-01','RR-MM-DD'),1,50,to_date('19-03-03','RR-MM-DD'),to_date('19-03-03','RR-MM-DD'),null,null,to_date('20-03-03','RR-MM-DD'),to_date('20-03-06','RR-MM-DD'),null);
Insert into ADMIN.PATIENTS (ID,PATIENT_ID,HOSPITAL_ID,SERIALID,TRIALGROUP,"INITIAL",BIRTHDATE,FEMALE,AGE,ALLOWDATE,STARTDATE,DROPDATE,MACCEDATE,DEADDATE,FINISHDATE,DIFFDAYS) values (8,'S2-002',8,2,1,'I.D',to_date('45-01-01','RR-MM-DD'),1,75,to_date('19-03-03','RR-MM-DD'),to_date('19-03-03','RR-MM-DD'),to_date('19-04-12','RR-MM-DD'),null,null,to_date('20-06-27','RR-MM-DD'),null);
--------------------------------------------------------
--  DDL for Index SYS_C0012044
--------------------------------------------------------

  CREATE UNIQUE INDEX "ADMIN"."SYS_C0012044" ON "ADMIN"."EVENTS" ("ID") 
  PCTFREE 10 INITRANS 2 MAXTRANS 255 COMPUTE STATISTICS 
  STORAGE(INITIAL 65536 NEXT 1048576 MINEXTENTS 1 MAXEXTENTS 2147483645
  PCTINCREASE 0 FREELISTS 1 FREELIST GROUPS 1
  BUFFER_POOL DEFAULT FLASH_CACHE DEFAULT CELL_FLASH_CACHE DEFAULT)
  TABLESPACE "DATA" ;
--------------------------------------------------------
--  DDL for Index SYS_C0012055
--------------------------------------------------------

  CREATE UNIQUE INDEX "ADMIN"."SYS_C0012055" ON "ADMIN"."HOSPITALS" ("ID") 
  PCTFREE 10 INITRANS 2 MAXTRANS 255 COMPUTE STATISTICS 
  STORAGE(INITIAL 65536 NEXT 1048576 MINEXTENTS 1 MAXEXTENTS 2147483645
  PCTINCREASE 0 FREELISTS 1 FREELIST GROUPS 1
  BUFFER_POOL DEFAULT FLASH_CACHE DEFAULT CELL_FLASH_CACHE DEFAULT)
  TABLESPACE "DATA" ;
--------------------------------------------------------
--  DDL for Index SYS_C0012061
--------------------------------------------------------

  CREATE UNIQUE INDEX "ADMIN"."SYS_C0012061" ON "ADMIN"."PATIENTS" ("ID") 
  PCTFREE 10 INITRANS 2 MAXTRANS 255 COMPUTE STATISTICS 
  STORAGE(INITIAL 65536 NEXT 1048576 MINEXTENTS 1 MAXEXTENTS 2147483645
  PCTINCREASE 0 FREELISTS 1 FREELIST GROUPS 1
  BUFFER_POOL DEFAULT FLASH_CACHE DEFAULT CELL_FLASH_CACHE DEFAULT)
  TABLESPACE "DATA" ;
--------------------------------------------------------
--  Constraints for Table EVENTS
--------------------------------------------------------

  ALTER TABLE "ADMIN"."EVENTS" MODIFY ("ID" NOT NULL ENABLE);
  ALTER TABLE "ADMIN"."EVENTS" ADD PRIMARY KEY ("ID")
  USING INDEX PCTFREE 10 INITRANS 2 MAXTRANS 255 COMPUTE STATISTICS 
  STORAGE(INITIAL 65536 NEXT 1048576 MINEXTENTS 1 MAXEXTENTS 2147483645
  PCTINCREASE 0 FREELISTS 1 FREELIST GROUPS 1
  BUFFER_POOL DEFAULT FLASH_CACHE DEFAULT CELL_FLASH_CACHE DEFAULT)
  TABLESPACE "DATA"  ENABLE;
--------------------------------------------------------
--  Constraints for Table HOSPITALS
--------------------------------------------------------

  ALTER TABLE "ADMIN"."HOSPITALS" MODIFY ("ID" NOT NULL ENABLE);
  ALTER TABLE "ADMIN"."HOSPITALS" MODIFY ("HOSPITAL_ID" NOT NULL ENABLE);
  ALTER TABLE "ADMIN"."HOSPITALS" MODIFY ("NAME" NOT NULL ENABLE);
  ALTER TABLE "ADMIN"."HOSPITALS" MODIFY ("CREATED_AT" NOT NULL ENABLE);
  ALTER TABLE "ADMIN"."HOSPITALS" MODIFY ("USERID" NOT NULL ENABLE);
  ALTER TABLE "ADMIN"."HOSPITALS" MODIFY ("USERPASS" NOT NULL ENABLE);
  ALTER TABLE "ADMIN"."HOSPITALS" ADD PRIMARY KEY ("ID")
  USING INDEX PCTFREE 10 INITRANS 2 MAXTRANS 255 COMPUTE STATISTICS 
  STORAGE(INITIAL 65536 NEXT 1048576 MINEXTENTS 1 MAXEXTENTS 2147483645
  PCTINCREASE 0 FREELISTS 1 FREELIST GROUPS 1
  BUFFER_POOL DEFAULT FLASH_CACHE DEFAULT CELL_FLASH_CACHE DEFAULT)
  TABLESPACE "DATA"  ENABLE;
--------------------------------------------------------
--  Constraints for Table PATIENTS
--------------------------------------------------------

  ALTER TABLE "ADMIN"."PATIENTS" MODIFY ("ID" NOT NULL ENABLE);
  ALTER TABLE "ADMIN"."PATIENTS" MODIFY ("PATIENT_ID" NOT NULL ENABLE);
  ALTER TABLE "ADMIN"."PATIENTS" ADD PRIMARY KEY ("ID")
  USING INDEX PCTFREE 10 INITRANS 2 MAXTRANS 255 COMPUTE STATISTICS 
  STORAGE(INITIAL 65536 NEXT 1048576 MINEXTENTS 1 MAXEXTENTS 2147483645
  PCTINCREASE 0 FREELISTS 1 FREELIST GROUPS 1
  BUFFER_POOL DEFAULT FLASH_CACHE DEFAULT CELL_FLASH_CACHE DEFAULT)
  TABLESPACE "DATA"  ENABLE;
