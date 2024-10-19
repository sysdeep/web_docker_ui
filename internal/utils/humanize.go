package utils

import "fmt"

const (
	ONE_KB = 1024
	//  ONE_KB_BI = BigInteger.valueOf(1024L);
	ONE_MB = 1048576
	//   public static final BigInteger ONE_MB_BI;
	ONE_GB = 1073741824
	//   public static final BigInteger ONE_GB_BI;
	ONE_TB = 1099511627776
	//public static final BigInteger ONE_TB_BI;
	ONE_PB = 1125899906842624
	//public static final BigInteger ONE_PB_BI;
	ONE_EB = 1152921504606846976
	//public static final BigInteger ONE_EB_BI;
	// public static final BigInteger ONE_ZB;
	// public static final BigInteger ONE_YB;
)

func ByteCountToDisplaySize(size int64) string {

	// try_eb := size / ONE_EB
	// if try_eb > 0 {
	// 	return fmt.Sprintf("%d EB", try_eb)
	// }

	// try_pb := size / ONE_PB
	// if try_pb > 0 {
	// 	return fmt.Sprintf("%d PB", try_pb)
	// }

	// try_tb := size / ONE_TB
	// if try_tb > 0 {
	// 	return fmt.Sprintf("%d TB", try_tb)
	// }

	try_gb := size / ONE_GB
	if try_gb > 0 {
		return fmt.Sprintf("%d GB", try_gb)
	}

	try_mb := size / ONE_MB
	if try_mb > 0 {
		return fmt.Sprintf("%d MB", try_mb)
	}

	try_kb := size / ONE_KB
	if try_kb > 0 {
		return fmt.Sprintf("%d KB", try_kb)
	}

	// Objects.requireNonNull(size, "size");
	// String displaySize;
	// if (size.divide(ONE_EB_BI).compareTo(BigInteger.ZERO) > 0) {
	//    displaySize = size.divide(ONE_EB_BI) + " EB";
	// } else if (size.divide(ONE_PB_BI).compareTo(BigInteger.ZERO) > 0) {
	//    displaySize = size.divide(ONE_PB_BI) + " PB";
	// } else if (size.divide(ONE_TB_BI).compareTo(BigInteger.ZERO) > 0) {
	//    displaySize = size.divide(ONE_TB_BI) + " TB";
	// } else if (size.divide(ONE_GB_BI).compareTo(BigInteger.ZERO) > 0) {
	//    displaySize = size.divide(ONE_GB_BI) + " GB";
	// } else if (size.divide(ONE_MB_BI).compareTo(BigInteger.ZERO) > 0) {
	//    displaySize = size.divide(ONE_MB_BI) + " MB";
	// } else if (size.divide(ONE_KB_BI).compareTo(BigInteger.ZERO) > 0) {
	//    displaySize = size.divide(ONE_KB_BI) + " KB";
	// } else {
	//    displaySize = size + " bytes";
	// }

	return fmt.Sprintf("%d bytes", size)
}
