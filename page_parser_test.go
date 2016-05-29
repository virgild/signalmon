package main

import "testing"

var diagnostics_page = `
<HTML>
<HEAD>
<TITLE>Thomson Cable Modem Technical Details Status</TITLE>
<META HTTP-EQUIV="Content-Type" CONTENT="text/html; charset=windows-1252">
<META HTTP-EQUIV="Refresh" CONTENT="7">
<link href="StyleA.css" type="text/css" rel="STYLESHEET">
<script type="text/javascript" language="javascript" src="menu.js"></script>
</head>
<body>
<table border="0" cellspacing="0" width="1003" id="AutoNumber1" height="488"  cellpadding="0">
  <TR>
    <TD width=238 valign="top">
        <TABLE WIDTH=238 BORDER=0 CELLPADDING=0 CELLSPACING=0  >
            <TR>
                <TD WIDTH=238 HEIGHT=98 COLSPAN=2 valign="top">
			        <IMG border=0 SRC="logo.jpg" ></TD>
            </TR>
            <TR>
        		<TD WIDTH=238 height="83" COLSPAN=2 valign="top">
			        <IMG border=0 SRC="u-04.jpg"></TD>
            </TR>
            <TR>
                <TD WIDTH=210 bgcolor="#dfdee2" height="420" align="right" valign="top">
                <br>
                    <TABLE WIDTH=180 HEIGHT=400 BORDER=0  CELLPADDING=0 CELLSPACING=0  >
                        <script type="text/javascript">
                        ShowSubMenu("Main", "Main.asp", "0");

                        ShowSubMenu("Diagnostics", "Diagnostics.asp", "1");

                        ShowSubMenu("Event Log", "EventLog.asp", "0");


                        </script>
                        <tr valign="bottom"><td class="footer"> &copy; - Thomson - 2007</td></tr>
                    </table>
                </TD>
        		<TD WIDTH=28 valign="top" background="u-09.jpg"><img src="u-09.jpg" width="28" height="32"></TD>
            </TR>
        </TABLE>
    </TD>
    <TD width=765 valign="top" >
        <TABLE WIDTH=765 BORDER=0 CELLPADDING=0 CELLSPACING=0  >
            <TR>
		    <TD HEIGHT=67 COLSPAN=3>
			    <IMG SRC="u-02.jpg" width="765" border=0 ></TD>
            </TR>
            <TR>
                <TD WIDTH=17 bgcolor="#dfdee2"> </TD>
       		  <TD height="31" width="720"  bgcolor="#dfdee2"> </TD>
        		<TD width="28" height="31" >
        			<IMG border=0 SRC="u-00.jpg" WIDTH=28 HEIGHT=31></TD>
        	</TR>
            <TR>
		        <TD COLSPAN=3>
			    <IMG WIDTH=765 HEIGHT=5 border=0 SRC="u-05.jpg" ></TD>
            </TR>

            <TR>
		        <TD COLSPAN=3>
                    <TABLE WIDTH=765 BORDER=0  CELLPADDING=0 CELLSPACING=0 >
                        <TR HEIGHT=60>
                            <TD>
                            <P class="desc"><b>MODEM TECHNICAL DETAILS STATUS PAGE</b>
                                <br><br> <small><small>This page will auto-refresh periodically.</small></small> </P></TD>
                        </TR>
                        <TR HEIGHT=400 align="center">
						  <TD valign="TOP" align="left"  >
							  <hr>
							  <br>
							  <p class="contents" >
                            <form action=/goform/Diagnostics method=POST name="Diagnostics">
                            <table class="light1">
                                   <tr>
                                  <td width="260">
                                  <font face="Arial"><small>Status Code:&nbsp;&nbsp;Operational</small></font>
                                  </td>
                                   </tr>
                               <tr>
                                  <td width="260" >
                                  <font  face="Arial"><small>Software Version:&nbsp;&nbsp;STAC.02.50</small></font>
                                  </td>
                               </tr>
                               <tr>
                                  <td width="260" >
                                  <font  face="Arial"><small>Software Model:&nbsp;&nbsp;a81a </small></font>
                                  </td>
                               </tr>
                               <tr>
                                  <td width="260" >
                                  <font  face="Arial"><small>Bootloader:&nbsp;&nbsp;2.3.1 </small></font>
                                  </td>
                               </tr>
                            </table>
                            <br>
                            <table class="light1">
                               <tr>
                                  <td>
                                  <small><strong><font face="Arial">Cable Signal Details</font></strong></small>
                                  </td>
                               </tr>
                               <tr>
                                  <td colspan="2" height="2">
                                  </td>
                               </tr>
                               <tr>
                                <td colspan=2>
                                        <font face="Arial"><small><strong>Forward Path: </strong></small><br /><table class="light" cellspacing=0 cellpadding=2 border=1><tr><td><b>Channel</b></td><td><b>Frequency</b></td><td><b>Power</b></td><td><b>SNR</b></td><td><b>BER</b></td><td><b>Modulation</b></td></tr><tr><td align="right">1</td><td align="right">693.0 MHz</td><td align="right">-1.6 dBmV</td><td align="right">40.3 dB</td><td align="right">0.000 %</td><td align="right">256 QAM</td></tr><tr><td align="right">2</td><td align="right">669.0 MHz</td><td align="right">-3.1 dBmV</td><td align="right">39.7 dB</td><td align="right">0.000 %</td><td align="right">256 QAM</td></tr><tr><td align="right">3</td><td align="right">675.0 MHz</td><td align="right">-2.7 dBmV</td><td align="right">39.9 dB</td><td align="right">0.000 %</td><td align="right">256 QAM</td></tr><tr><td align="right">4</td><td align="right">681.0 MHz</td><td align="right">-2.9 dBmV</td><td align="right">39.7 dB</td><td align="right">0.000 %</td><td align="right">256 QAM</td></tr><tr><td align="right">5</td><td align="right">687.0 MHz</td><td align="right">-2.1 dBmV</td><td align="right">40.2 dB</td><td align="right">0.000 %</td><td align="right">256 QAM</td></tr><tr><td align="right">6</td><td align="right">699.0 MHz</td><td align="right">-1.2 dBmV</td><td align="right">40.3 dB</td><td align="right">0.000 %</td><td align="right">256 QAM</td></tr><tr><td align="right">7</td><td align="right">705.0 MHz</td><td align="right">-1.0 dBmV</td><td align="right">40.3 dB</td><td align="right">0.000 %</td><td align="right">256 QAM</td></tr><tr><td align="right">8</td><td align="right">711.0 MHz</td><td align="right">-1.0 dBmV</td><td align="right">40.2 dB</td><td align="right">0.049 %</td><td align="right">256 QAM</td></tr></table>
                                </td>
                               </tr>
                               <tr>
                                  <td colspan="2">
                                  </td>
                               </tr>
                               <tr>

                                  <td colspan=2>
                                        <font face="Arial"><small><strong>Return Path: </strong></small><br /><table class="light" cellspacing=0 cellpadding=2 border=1><tr><td><b>Channel ID</b></td><td><b>Frequency</b></td><td><b>Power</b></td><td><b>Modulation</b></td></tr><tr><td align="right">3</td><td align="right">23.7 MHz</td><td align="right">45.0 dBmV</td><td align="right">64 QAM</td></tr><tr><td align="right">1</td><td align="right">38.6 MHz</td><td align="right">45.0 dBmV</td><td align="right">64 QAM</td></tr><tr><td align="right">2</td><td align="right">30.6 MHz</td><td align="right">45.0 dBmV</td><td align="right">64 QAM</td></tr><tr><td align="right">0</td><td align="right">0.0 MHz</td><td align="right">0.0 dBmV</td><td align="right">QPSK</td></tr></table>
                                  </td>
                               </tr>
                               <tr>
                                  <td colspan="2" height="2">
                                  </td>
                               </tr>
                               <tr>
                                  <td >
                                  <small><strong><font face="Arial">Data Service Details</font></strong></small>
                                  </td>
                               </tr>
                               <tr>
                                <td>
                                  <table width=290>
                                  <tr>
                                    <td width=30></td>
                                    <td>
                                        <font face="Arial"><small>Provisioned Address: Yes</small></font>
                                    </td>
                                  </tr>
                                  <tr>
                                    <td width=30></td>
                                    <td>
                                        <font face="Arial"><small>Provisioned Time: Yes</small></font>
                                    </td>
                                  </tr>
                                  <tr>
                                    <td width=30></td>
                                    <td>
                                        <font face="Arial"><small>Provisioned Configuration: Yes</small></font>
                                    </td>
                                  </tr>
                                  <tr>
                                    <td width=30></td>
                                    <td>
                                        <font face="Arial"><small>Registered: Yes</small></font>
                                    </td>
                                  </tr>
                                  <tr>
                                    <td width=30></td>
                                    <td>
                                        <font face="Arial"><small>BPI: Enabled</small></font>
                                    </td>
                                  </tr>
                                  </table>
                                </td>
                               </tr>
                            </table>
                            </form> </p>
                            </td>
						  </TR>
                  </TABLE>
                </TD>
            </TR>
      </TABLE>
    </TD>
</TR>
</TABLE>
</BODY>
</HTML>
`

func TestParseDiagnosticsPage(t *testing.T) {
	statsdata, err := ParseDiagnosticsPage(diagnostics_page)
	if err != nil {
		t.Errorf("ParsePage() return an error: %v", err)
	}

	if statsdata.ForwardSignals[1].Power() != -3.1 {
		t.Errorf("statsdata.ForwardSignals[1].Power() != -3.1")
	}

	if statsdata.ForwardSignals[1].Frequency() != 669.0 {
		t.Errorf("statsdata.ForwardSignals[1].Frequency() != 669.0 (%f)", statsdata.ForwardSignals[1].Frequency())
	}
}
