package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/disburse_db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Exec(
		"CREATE TABLE IF NOT EXISTS `disburses` (" +
			"`id` int(11) unsigned NOT NULL AUTO_INCREMENT," +
			"`transaction_id` bigint(20) DEFAULT NULL," +
			"`amount` int(11) DEFAULT NULL," +
			"`status` varchar(10) DEFAULT ''," +
			"`timestamp` timestamp NULL DEFAULT NULL," +
			"`bank_code` varchar(10) DEFAULT NULL," +
			"`account_number` varchar(20) DEFAULT NULL," +
			"`beneficiary_name` varchar(50) DEFAULT NULL," +
			"`remark` text," +
			"`receipt` varchar(255) DEFAULT ''," +
			"`time_served` timestamp NULL DEFAULT NULL," +
			"`fee` int(11) DEFAULT '0'," +
			"PRIMARY KEY (`id`)" +
			") ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;",
	)

	db.Exec(
		"INSERT INTO `disburses` (`id`, `transaction_id`, `amount`, `status`, `timestamp`, `bank_code`, `account_number`, `beneficiary_name`, `remark`, `receipt`, `time_served`, `fee`)" +
			"VALUES" +
			"(1, 1879754755, 10000, 'SUCCESS', '2020-05-08 23:09:42', 'bni', '1234567890', 'PT FLIP', 'sample remark', 'https://flip-receipt.oss-ap-southeast-5.aliyuncs.com/debit_receipt/126316_3d07f9fef9612c7275b3c36f7e1e5762.jpg', '2020-05-08 23:18:42', 4000)," +
			"(2, 1402967034, 1000, 'PENDING', '2020-05-09 13:06:41', '303', '12345678', 'PT FLIP', 'remark', NULL, NULL, 4000)," +
			"(3, 6675673941, 1000, 'PENDING', '2020-05-09 13:16:48', '303', '12345678', 'PT FLIP', 'remark', NULL, NULL, 4000)," +
			"(4, 1202491636, 1000, 'PENDING', '2020-05-09 13:19:21', '303', '12345678', 'PT FLIP', 'remark', NULL, NULL, 4000)," +
			"(5, 1246224051, 10000, 'SUCCESS', '2020-05-09 13:15:09', 'bni', '1234567890', 'PT FLIP', 'sample remark', 'https://flip-receipt.oss-ap-southeast-5.aliyuncs.com/debit_receipt/126316_3d07f9fef9612c7275b3c36f7e1e5762.jpg', '2020-05-09 13:24:09', 4000)," +
			"(6, 7649561267, 1000, 'SUCCESS', '2020-05-09 13:27:33', '303', '12345678', 'PT FLIP', 'remark', 'https://flip-receipt.oss-ap-southeast-5.aliyuncs.com/debit_receipt/126316_3d07f9fef9612c7275b3c36f7e1e5762.jpg', '2020-05-09 13:26:52', 4000)," +
			"(7, 8317074329, 10000, 'SUCCESS', '2020-05-09 13:42:23', 'bni', '1234567890', 'PT FLIP', 'sample remark', 'https://flip-receipt.oss-ap-southeast-5.aliyuncs.com/debit_receipt/126316_3d07f9fef9612c7275b3c36f7e1e5762.jpg', '2020-05-09 14:06:44', 4000);",
	)

}
