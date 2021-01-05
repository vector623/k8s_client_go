package k8s_client_go

import "testing"

func TestGenerateKubeConf(t *testing.T) {
	type args struct {
		name string
		serverUrl             string
		clientCertificateData string
		clientKeyData         string
		token                 string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "clusterUser_feiseu2-supply-rg-001_tfprovtest",
			args: args{
				name: "clusterUser_feiseu2-supply-rg-001_tfprovtest",
				serverUrl:             "https://tfprovtest-c6a942aa.hcp.eastus2.azmk8s.io:443",
				clientCertificateData: "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUUvRENDQXVTZ0F3SUJBZ0lRSWRYQkdqKzlBbi8vb1lIZUs4RTRiekFOQmdrcWhraUc5dzBCQVFzRkFEQU4KTVFzd0NRWURWUVFERXdKallUQWVGdzB5TVRBeE1EUXhPRFUzTkRKYUZ3MHlNekF4TURReE9UQTNOREphTURBeApGekFWQmdOVkJBb1REbk41YzNSbGJUcHRZWE4wWlhKek1SVXdFd1lEVlFRREV3eHRZWE4wWlhKamJHbGxiblF3CmdnSWlNQTBHQ1NxR1NJYjNEUUVCQVFVQUE0SUNEd0F3Z2dJS0FvSUNBUUNpcFVubkY5L3FIZUEzbFI4cStkQSsKV0xRMUVvUXlhc3g0NllQMkltMnlacVlVTVVRaERnODNzUHZGRjVlSCtiNlRRc3VuUTFJTlVoUnFjc1NuODN2cgpJOFYweVZlUUJPWEhxOThtQ21TMFZEM1FoaU1CdU0rTUVuL3VOckNaZHdGT2xmSjU1TXI1Mm1GTDZhK2gwYjdJCkpQSFJERnJJeTk0bld6Vi9VTUczK0lLdk1ZdXgwUUpOQnlsak0wUG9TL2tOblpPeUxXRVdBRkcwZnN2aUdVWm8Kd3czUlN3MWxPV1BxT1p4NStmc0xMZFBTcE1vUyt5N1NwVWpHdytnazBJZ2hZZ0l6WC9JZUxKWmJXZmFjOFlPUApITHB5OGwxTW9MeWdzaXdla3M3UUlzZmN5RzErblVVTHNHV2dUUk5mWnVySUtRbTllT2dURzd6WGdqenpHUytKCk1aWGpFa1BVVTRIU0V4WkVTTWFSU0xzVWZTcnkzRFplaG1yUFhqc1Q4Tm5QTEt5Z1JiMVRma2plaTlURDU0WU8Kb0w4MWQzdVlDS1JLUis3anRQRUtpNDAvWU5nTjNtdnd2YmhodzRYakdwaW4rOWFrSTN2MmdKQ2NIeS95b05nTwpIMjZLUTBiSzBZZm03WlJES2RSSncrYVg5T08wdERaS3RqYm43UVkxeDkrdVp2NlVIeG1OZ2RUMnhDMDhadk5qCnZteklHN0N6WEQ2Yk9DVjZKVmxoNjRsblh6NUlJUys1QzZ1dm01QVkvUUlBNm90a2NnRzNMa0lmdlJsKzYrbmcKRHU4TGlYMFFsdVF3RHJpZElVUXZnSzdQd1c3WDdrWWMvOEt1c0VHVTcyN1ZiNXNvZVhkUGFWdFRsRGJYaHY3aApWSk1tM0krdFV6aDNwNTVCbWZ3dENRSURBUUFCb3pVd016QU9CZ05WSFE4QkFmOEVCQU1DQmFBd0V3WURWUjBsCkJBd3dDZ1lJS3dZQkJRVUhBd0l3REFZRFZSMFRBUUgvQkFJd0FEQU5CZ2txaGtpRzl3MEJBUXNGQUFPQ0FnRUEKb3lDVkFDN2c4MXdacWtUY0xYMmJlcnhCRjZnbXBRejZCaHBESW5LNGxINGhFaFZqMU81MVFralVMY0twN29uagptOGIwSTlWNkxvMDFXckpjaUJFZC9wS2VCMEhLQXMwaGlvTHNWZDFieUI5ZU5WNE9xcFhWd3ViVlBZSlpSVlVNCjlHTWdpenJnbkFWQ25GN1pFc3lIWi9GZ1g0TFkvT0pQZnlTejdMQjBpNzhnbjVaOU81Nm9zWEYvT1JlLzdKWVAKbkdxekFJSEoyV3pKSmZHdW9VaTA2bUttSEUrclJXUm5nOElpYTZ5Z3A5ME9LcWZQRHNvMjIrNWpIVG9hMFBzdQpHRG5QOTBCa0lUeWZKNFcyUEZnblpMclJNM25ZTXNZczRYV29GdkU0aFdjaDBHSi9UWWpUdFdZQnBiMmtnU2I4CmNXYVk3OUdxV3JnRmFYUjcxbkZ0SE92ZzRzbFBUZy80WU5GOUw4UTZiVWRLOVF6VjZPK2xFQ2pCN0RvcmZrS2oKaE9rMU95dnFmQVIzUjM2RUtobitRTTFDbkx3RFIydEE1cm9tcG9JbDBoeFB1QlZWa05ndkRWQnYwODFCSG0vTAoxNkpRWm13MGRpTmlobFhUcXlxeVc5bDFkYWk0bXlGWUt5dk9vd3lMU04wc054WG1qVlBaY3o2ZjZOSjEyMGpFCkhMelBNYVVDY3ZqeGRXT2RnVnpLRmNEam5JYkVCbDI1SFoyMHE5Ujlva1FHbmJCNmFtNDdCV3pqUkFwR01Rc3gKRjFKZ1JkVncxMXJvNERxcGhOTkJKNzlCbGphVC9obndhMVhBMGdCRUxmcWxVeTliWDlTMzQ5WktyWHAwZlBvMQpEaFpXK3F6Q0t6LzhNQ25hUHcrM2p1UGdJQ3dDaHFmN291NFlocUNIa2pFPQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==",
				clientKeyData:         "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlKS0FJQkFBS0NBZ0VBb3FWSjV4ZmY2aDNnTjVVZkt2blFQbGkwTlJLRU1tck1lT21EOWlKdHNtYW1GREZFCklRNFBON0Q3eFJlWGgvbStrMExMcDBOU0RWSVVhbkxFcC9ONzZ5UEZkTWxYa0FUbHg2dmZKZ3BrdEZROTBJWWoKQWJqUGpCSi83amF3bVhjQlRwWHllZVRLK2RwaFMrbXZvZEcreUNUeDBReGF5TXZlSjFzMWYxREJ0L2lDcnpHTApzZEVDVFFjcFl6TkQ2RXY1RFoyVHNpMWhGZ0JSdEg3TDRobEdhTU1OMFVzTlpUbGo2am1jZWZuN0N5M1QwcVRLCkV2c3UwcVZJeHNQb0pOQ0lJV0lDTTEveUhpeVdXMW4yblBHRGp4eTZjdkpkVEtDOG9MSXNIcExPMENMSDNNaHQKZnAxRkM3QmxvRTBUWDJicXlDa0p2WGpvRXh1ODE0STg4eGt2aVRHVjR4SkQxRk9CMGhNV1JFakdrVWk3RkgwcQo4dHcyWG9acXoxNDdFL0Raenl5c29FVzlVMzVJM292VXcrZUdEcUMvTlhkN21BaWtTa2Z1NDdUeENvdU5QMkRZCkRkNXI4TDI0WWNPRjR4cVlwL3ZXcENONzlvQ1FuQjh2OHFEWURoOXVpa05HeXRHSDV1MlVReW5VU2NQbWwvVGoKdExRMlNyWTI1KzBHTmNmZnJtYitsQjhaallIVTlzUXRQR2J6WTc1c3lCdXdzMXcrbXpnbGVpVlpZZXVKWjE4KwpTQ0V2dVF1cnI1dVFHUDBDQU9xTFpISUJ0eTVDSDcwWmZ1dnA0QTd2QzRsOUVKYmtNQTY0blNGRUw0Q3V6OEZ1CjErNUdIUC9DcnJCQmxPOXUxVytiS0hsM1QybGJVNVEyMTRiKzRWU1RKdHlQclZNNGQ2ZWVRWm44TFFrQ0F3RUEKQVFLQ0FnQnEwUVQ5S1hmY0NPejFkQkU1SjBHMWdJb3hoTXlmVS9hZnZqVDFWRjBZMmtOZVpaQ3A2TkxqcjI5YQpYSWtmTzAwRzU4cGFPbU5MN2trTU0vMmxrOG9oZnFleWNIejZUODBjMm1xK1Qvc2RWMm9WWmRMZDBhTTFDZjFNCjFFd2RMY1FFK1lONnZHRDF3STgvek9TVThPeE1BNzl3bkoyTzNCMDFXYmlEdWVUUHp2ejgxQnc3bHRuYy9zUVcKcTFlQTIweGd2Y2wvTGtCaHM2bjFVWVRUSStoUzZ5OUs4ZkNGTnpEQVpqKzIrQThBbzROSEJ3VFRmSFdlMEFCQQowZTdISnRFUW80NXdjU1hRM1ZXTi8wSndoa2p4bDF0NEhHWmI2Z0NYeFF3bXRDSHhIYmFHOEhiQXZjYlh5MnlpCjM1dkdmdy9ZanVNZnRhVXN6MzJ4UjQ2QjFWdWF3WVVteExvZmh0L0hxd0c2TU1nT0cyeUhlVThQQ3oxTWNka0YKWkI3N296UDI2TjRiOUJOWEY5dkN4dUVaWTNwU01QRnRjY0JaNnFXREZORHpxUzRONTZpRmxOMDVqOTdGb3pJUgpEMnUvalpTYjJBOUYwWVlkQ3pxdW1GbEVWYUNLbCtqdFpIeENCQlFxdStMRkxLNnY1Q3Q3NVFXbUw2WUhXVnlCCnptYWVzQU9RZm93d0JsaWFOWVZ6VXdwakEyZGJVUGJhQlB1M2xNYkNPcHZhQUxIbklZU3FtSGdNSmtwWjNSa0MKd0lyODhYKzFXNmE0czJtcHhGWkNTN0YxbzQwU1BZaGN3dlZrRGY3Qkt6aDl6elZhMytLaUhqNitVaVVVL2RMSwphM0xDbHdyeHpxUUdlRmhSYTAyVEFrZTYyVndZSnhMcm5Rc2VTbVZQRStqTUg3cEhsUUtDQVFFQTBRQ1JGTTkrCk01MHg5R3B5Z3NFOHBjU1J0UlRwS201UTZSdHhNMnlGV3JvZ3M0V1VqV2UreGVIdXhsUnRtSTd1QkViMWg3aTkKcE8rMmVyeTkvRmF1UHpvaDZwKzJqckMxZU9HMytISEVYOXd1dW02diswZ01MY3djMVRYSjRhRWIrbEZuazdyaQoxNkNGZEJMUjJJVnFGK29FcVY5QVl2bEVXb0o1Z0JWQTgvdEd5bEdBcm5pYXVUWUV6UVhsQkZxMGdpUnhvRmhICm5ZcDhQbHpsYU1wV3pkU2RvTk9XUGFOdjBTd2o3K2xJQkwwbUtOajA5elFISDlnUUtaVytRSWY2UjVQOWZMczYKMUQvRmI4d3BSRFh4T3FVS0paTEZSbVo3Z3U0UWNlRWY2UndEekZIQjFHQ2pWSHdiV2ZFRXoxeHl5T1lJOFNOTAppelVoUFJkNHhnamJGd0tDQVFFQXh6Z29hK1p0Q3FJOUdJcmxwMitKRm1nQitnUE12VXpqUGNxUVJWSHNBZmRMCmtXaFpFZ0I0L1l2YzRpUUxFN2lLWEtETXR1SWxncXJMZUo0amwwd0E1TS93eElHanNWaFBjeldmNUtsS25FaE4KNTJFVUsxNFFhc00velVJYkdPb1dQVllQYkpzY216NUVnS1BZTHJMVjFEdTNqa3BSdElreFRocGV0OGs1Y3lsTwpvdTlnK1BkL2FjTGI2SFBQNXF2a1JTOTUzY29PZCtTUFc0ZUxQeDBwWHIybjFCeFBsTkhZU0dHVTFUTGNWdEZLCkRsV0wyeDdJL3RsRTlVclM5d0s0eXJOaVdLbU15OUNlby9zM0ExN1kreGU0TW4wc0RiaTVyeFE0LzUxRitRSlYKZkU1RC95UElWNFJzb254cTJCZVNMeFlNMU80cFRjQ0k2V3NBaFNuTTN3S0NBUUVBazIxMzZac1g1YWFyOTBJagpObkxDcWhnSHhCbFZCRFkrNHNBVjBQVnE5LzIxVlkwVHBjK1p0WkpTMFJuN0dSTXhlRmpwbDJ4TWNnMFp4bWhmCkI2SEpWaFpXbUY5QzhsSmFJRVB3ZHRNaGQ1RWtjU1BsQlRia1NHWGpyTTQ5Y3VnbmlTSmxFc05RZ2xQYjcyM2IKRW9YQnVjNjhyUkRncFRBNTlESjNMRTlSdXdqei9ZRVNhWC92ZkZoTGVLZ0F6VnFDWEE4REM5MVozeG94Mi9mQgpaM0xiV21QeE1PSXY1NWpqSmtTaHNQajRjVWRQcEo2dDViYVFrL2pnSmNWWUo4TURXL3hiWWVIY2RKM3JVcHBNClBhNWpWRC9ZOFhvZnFESmdTNUpNalZnRkgzVkFoSjdqUUVGTzFTaWFUbzZDR3hlV2JIby9zb0Y0SjJFNTZoTWwKZ1dQN3p3S0NBUUE3dDRTK1cxTWlWZUdVZ3NZMmU0T1BXTndHMHE3RHE5R09Tek1WNUdtbmFUYVl2cnU4WHJIbAo1cy91WmJtQVR0Z0N6aHl2VVhqSEEyVjhvWUt2cnBPeUF2YjhJSlF0Ni9mYzVCMHEwT2hSa29idnJiVDdEdlpDCjJvWThnKzFWZVkzekV1SUt2ekJEQk5aSjdWTTNKRnBlaFF3UHRnMXAzWm9PL0c2YUR0S094N29QczB2Uk5YZFcKLzVjbHhkbHpIdUdOTDU4TE1hOUdseGUxNnhwOWI4TmY1OS93QUJDRzNxUkRoaVMwQjh2dXBrdlpvaUVMTE40NQpMbTJ5VFp5UmVKcG1BSUdoYkNtTlN0clM4dEhaeTh1azhNWkJOQzVKYXE4V1lTeDR5dThYcG5Cems0NWtZNWlCClMydTVCU3lOMi9ydzdGNHFVYjNGMWNuMGxpT1F4eWZuQW9JQkFGTlE0am1uWTZLQmtvVUZPTXE3TUtpUHEyM1kKMEhMdGduNTVTTlk1dUNuNmxaSGtJSTZVOTlUSWNqN24xanZFV2ZkL09oNS95T1RNWGd6VndFMUNFTXpNSVVFRgpqYlhYcWFHR3JnK1lWbFA4Vk5NdDgrc1hqRTZ4RHJVMFJTWEY0MVlySTljSVJZS0FRR3BYa0x4SzdCUHhnZm1DCm1yV2NUZVZrQlJpV1UzejVEQUJiSVV5emxpV09HNFIxZWxYQlUrWVVGOUlmUW1GcWRkSU1oVGQ3UGZWWjNOYnUKd2RxV3VGM1N1bXRzK3VBQ2VXTmJ5MVR4cnBScXMxVDMvdzNSM1VPaTJ3K0E3M1lSU2xiOEtwWUUybkEyamRZaApXeGRyMnNldVN6bjZ6dVVsTytJVDBJREVNakx6c3FDSTRzQ3JQak4rcm44U1dZVC8zVWNCRVRzY2ZFcz0KLS0tLS1FTkQgUlNBIFBSSVZBVEUgS0VZLS0tLS0K",
				token:                 "93c48444c37c36b9f37fb8962591decfe3cf0507d0a4dc55beb089074aa97863330676baeb260e134bf7c338671cc4562ca6c87bc02109dad2672ef3b58abf55",
			},
			//want: AksCreds{
			//	Name: "clusterUser_feiseu2-supply-rg-001_tfprovtest",
			//	User: struct {
			//		ClientCertificateData string
			//		ClientKeyData         string
			//		Token                   string
			//	}{
			//		ClientCertificateData: "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUUvRENDQXVTZ0F3SUJBZ0lRSWRYQkdqKzlBbi8vb1lIZUs4RTRiekFOQmdrcWhraUc5dzBCQVFzRkFEQU4KTVFzd0NRWURWUVFERXdKallUQWVGdzB5TVRBeE1EUXhPRFUzTkRKYUZ3MHlNekF4TURReE9UQTNOREphTURBeApGekFWQmdOVkJBb1REbk41YzNSbGJUcHRZWE4wWlhKek1SVXdFd1lEVlFRREV3eHRZWE4wWlhKamJHbGxiblF3CmdnSWlNQTBHQ1NxR1NJYjNEUUVCQVFVQUE0SUNEd0F3Z2dJS0FvSUNBUUNpcFVubkY5L3FIZUEzbFI4cStkQSsKV0xRMUVvUXlhc3g0NllQMkltMnlacVlVTVVRaERnODNzUHZGRjVlSCtiNlRRc3VuUTFJTlVoUnFjc1NuODN2cgpJOFYweVZlUUJPWEhxOThtQ21TMFZEM1FoaU1CdU0rTUVuL3VOckNaZHdGT2xmSjU1TXI1Mm1GTDZhK2gwYjdJCkpQSFJERnJJeTk0bld6Vi9VTUczK0lLdk1ZdXgwUUpOQnlsak0wUG9TL2tOblpPeUxXRVdBRkcwZnN2aUdVWm8Kd3czUlN3MWxPV1BxT1p4NStmc0xMZFBTcE1vUyt5N1NwVWpHdytnazBJZ2hZZ0l6WC9JZUxKWmJXZmFjOFlPUApITHB5OGwxTW9MeWdzaXdla3M3UUlzZmN5RzErblVVTHNHV2dUUk5mWnVySUtRbTllT2dURzd6WGdqenpHUytKCk1aWGpFa1BVVTRIU0V4WkVTTWFSU0xzVWZTcnkzRFplaG1yUFhqc1Q4Tm5QTEt5Z1JiMVRma2plaTlURDU0WU8Kb0w4MWQzdVlDS1JLUis3anRQRUtpNDAvWU5nTjNtdnd2YmhodzRYakdwaW4rOWFrSTN2MmdKQ2NIeS95b05nTwpIMjZLUTBiSzBZZm03WlJES2RSSncrYVg5T08wdERaS3RqYm43UVkxeDkrdVp2NlVIeG1OZ2RUMnhDMDhadk5qCnZteklHN0N6WEQ2Yk9DVjZKVmxoNjRsblh6NUlJUys1QzZ1dm01QVkvUUlBNm90a2NnRzNMa0lmdlJsKzYrbmcKRHU4TGlYMFFsdVF3RHJpZElVUXZnSzdQd1c3WDdrWWMvOEt1c0VHVTcyN1ZiNXNvZVhkUGFWdFRsRGJYaHY3aApWSk1tM0krdFV6aDNwNTVCbWZ3dENRSURBUUFCb3pVd016QU9CZ05WSFE4QkFmOEVCQU1DQmFBd0V3WURWUjBsCkJBd3dDZ1lJS3dZQkJRVUhBd0l3REFZRFZSMFRBUUgvQkFJd0FEQU5CZ2txaGtpRzl3MEJBUXNGQUFPQ0FnRUEKb3lDVkFDN2c4MXdacWtUY0xYMmJlcnhCRjZnbXBRejZCaHBESW5LNGxINGhFaFZqMU81MVFralVMY0twN29uagptOGIwSTlWNkxvMDFXckpjaUJFZC9wS2VCMEhLQXMwaGlvTHNWZDFieUI5ZU5WNE9xcFhWd3ViVlBZSlpSVlVNCjlHTWdpenJnbkFWQ25GN1pFc3lIWi9GZ1g0TFkvT0pQZnlTejdMQjBpNzhnbjVaOU81Nm9zWEYvT1JlLzdKWVAKbkdxekFJSEoyV3pKSmZHdW9VaTA2bUttSEUrclJXUm5nOElpYTZ5Z3A5ME9LcWZQRHNvMjIrNWpIVG9hMFBzdQpHRG5QOTBCa0lUeWZKNFcyUEZnblpMclJNM25ZTXNZczRYV29GdkU0aFdjaDBHSi9UWWpUdFdZQnBiMmtnU2I4CmNXYVk3OUdxV3JnRmFYUjcxbkZ0SE92ZzRzbFBUZy80WU5GOUw4UTZiVWRLOVF6VjZPK2xFQ2pCN0RvcmZrS2oKaE9rMU95dnFmQVIzUjM2RUtobitRTTFDbkx3RFIydEE1cm9tcG9JbDBoeFB1QlZWa05ndkRWQnYwODFCSG0vTAoxNkpRWm13MGRpTmlobFhUcXlxeVc5bDFkYWk0bXlGWUt5dk9vd3lMU04wc054WG1qVlBaY3o2ZjZOSjEyMGpFCkhMelBNYVVDY3ZqeGRXT2RnVnpLRmNEam5JYkVCbDI1SFoyMHE5Ujlva1FHbmJCNmFtNDdCV3pqUkFwR01Rc3gKRjFKZ1JkVncxMXJvNERxcGhOTkJKNzlCbGphVC9obndhMVhBMGdCRUxmcWxVeTliWDlTMzQ5WktyWHAwZlBvMQpEaFpXK3F6Q0t6LzhNQ25hUHcrM2p1UGdJQ3dDaHFmN291NFlocUNIa2pFPQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==",
			//		ClientKeyData:         "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlKS0FJQkFBS0NBZ0VBb3FWSjV4ZmY2aDNnTjVVZkt2blFQbGkwTlJLRU1tck1lT21EOWlKdHNtYW1GREZFCklRNFBON0Q3eFJlWGgvbStrMExMcDBOU0RWSVVhbkxFcC9ONzZ5UEZkTWxYa0FUbHg2dmZKZ3BrdEZROTBJWWoKQWJqUGpCSi83amF3bVhjQlRwWHllZVRLK2RwaFMrbXZvZEcreUNUeDBReGF5TXZlSjFzMWYxREJ0L2lDcnpHTApzZEVDVFFjcFl6TkQ2RXY1RFoyVHNpMWhGZ0JSdEg3TDRobEdhTU1OMFVzTlpUbGo2am1jZWZuN0N5M1QwcVRLCkV2c3UwcVZJeHNQb0pOQ0lJV0lDTTEveUhpeVdXMW4yblBHRGp4eTZjdkpkVEtDOG9MSXNIcExPMENMSDNNaHQKZnAxRkM3QmxvRTBUWDJicXlDa0p2WGpvRXh1ODE0STg4eGt2aVRHVjR4SkQxRk9CMGhNV1JFakdrVWk3RkgwcQo4dHcyWG9acXoxNDdFL0Raenl5c29FVzlVMzVJM292VXcrZUdEcUMvTlhkN21BaWtTa2Z1NDdUeENvdU5QMkRZCkRkNXI4TDI0WWNPRjR4cVlwL3ZXcENONzlvQ1FuQjh2OHFEWURoOXVpa05HeXRHSDV1MlVReW5VU2NQbWwvVGoKdExRMlNyWTI1KzBHTmNmZnJtYitsQjhaallIVTlzUXRQR2J6WTc1c3lCdXdzMXcrbXpnbGVpVlpZZXVKWjE4KwpTQ0V2dVF1cnI1dVFHUDBDQU9xTFpISUJ0eTVDSDcwWmZ1dnA0QTd2QzRsOUVKYmtNQTY0blNGRUw0Q3V6OEZ1CjErNUdIUC9DcnJCQmxPOXUxVytiS0hsM1QybGJVNVEyMTRiKzRWU1RKdHlQclZNNGQ2ZWVRWm44TFFrQ0F3RUEKQVFLQ0FnQnEwUVQ5S1hmY0NPejFkQkU1SjBHMWdJb3hoTXlmVS9hZnZqVDFWRjBZMmtOZVpaQ3A2TkxqcjI5YQpYSWtmTzAwRzU4cGFPbU5MN2trTU0vMmxrOG9oZnFleWNIejZUODBjMm1xK1Qvc2RWMm9WWmRMZDBhTTFDZjFNCjFFd2RMY1FFK1lONnZHRDF3STgvek9TVThPeE1BNzl3bkoyTzNCMDFXYmlEdWVUUHp2ejgxQnc3bHRuYy9zUVcKcTFlQTIweGd2Y2wvTGtCaHM2bjFVWVRUSStoUzZ5OUs4ZkNGTnpEQVpqKzIrQThBbzROSEJ3VFRmSFdlMEFCQQowZTdISnRFUW80NXdjU1hRM1ZXTi8wSndoa2p4bDF0NEhHWmI2Z0NYeFF3bXRDSHhIYmFHOEhiQXZjYlh5MnlpCjM1dkdmdy9ZanVNZnRhVXN6MzJ4UjQ2QjFWdWF3WVVteExvZmh0L0hxd0c2TU1nT0cyeUhlVThQQ3oxTWNka0YKWkI3N296UDI2TjRiOUJOWEY5dkN4dUVaWTNwU01QRnRjY0JaNnFXREZORHpxUzRONTZpRmxOMDVqOTdGb3pJUgpEMnUvalpTYjJBOUYwWVlkQ3pxdW1GbEVWYUNLbCtqdFpIeENCQlFxdStMRkxLNnY1Q3Q3NVFXbUw2WUhXVnlCCnptYWVzQU9RZm93d0JsaWFOWVZ6VXdwakEyZGJVUGJhQlB1M2xNYkNPcHZhQUxIbklZU3FtSGdNSmtwWjNSa0MKd0lyODhYKzFXNmE0czJtcHhGWkNTN0YxbzQwU1BZaGN3dlZrRGY3Qkt6aDl6elZhMytLaUhqNitVaVVVL2RMSwphM0xDbHdyeHpxUUdlRmhSYTAyVEFrZTYyVndZSnhMcm5Rc2VTbVZQRStqTUg3cEhsUUtDQVFFQTBRQ1JGTTkrCk01MHg5R3B5Z3NFOHBjU1J0UlRwS201UTZSdHhNMnlGV3JvZ3M0V1VqV2UreGVIdXhsUnRtSTd1QkViMWg3aTkKcE8rMmVyeTkvRmF1UHpvaDZwKzJqckMxZU9HMytISEVYOXd1dW02diswZ01MY3djMVRYSjRhRWIrbEZuazdyaQoxNkNGZEJMUjJJVnFGK29FcVY5QVl2bEVXb0o1Z0JWQTgvdEd5bEdBcm5pYXVUWUV6UVhsQkZxMGdpUnhvRmhICm5ZcDhQbHpsYU1wV3pkU2RvTk9XUGFOdjBTd2o3K2xJQkwwbUtOajA5elFISDlnUUtaVytRSWY2UjVQOWZMczYKMUQvRmI4d3BSRFh4T3FVS0paTEZSbVo3Z3U0UWNlRWY2UndEekZIQjFHQ2pWSHdiV2ZFRXoxeHl5T1lJOFNOTAppelVoUFJkNHhnamJGd0tDQVFFQXh6Z29hK1p0Q3FJOUdJcmxwMitKRm1nQitnUE12VXpqUGNxUVJWSHNBZmRMCmtXaFpFZ0I0L1l2YzRpUUxFN2lLWEtETXR1SWxncXJMZUo0amwwd0E1TS93eElHanNWaFBjeldmNUtsS25FaE4KNTJFVUsxNFFhc00velVJYkdPb1dQVllQYkpzY216NUVnS1BZTHJMVjFEdTNqa3BSdElreFRocGV0OGs1Y3lsTwpvdTlnK1BkL2FjTGI2SFBQNXF2a1JTOTUzY29PZCtTUFc0ZUxQeDBwWHIybjFCeFBsTkhZU0dHVTFUTGNWdEZLCkRsV0wyeDdJL3RsRTlVclM5d0s0eXJOaVdLbU15OUNlby9zM0ExN1kreGU0TW4wc0RiaTVyeFE0LzUxRitRSlYKZkU1RC95UElWNFJzb254cTJCZVNMeFlNMU80cFRjQ0k2V3NBaFNuTTN3S0NBUUVBazIxMzZac1g1YWFyOTBJagpObkxDcWhnSHhCbFZCRFkrNHNBVjBQVnE5LzIxVlkwVHBjK1p0WkpTMFJuN0dSTXhlRmpwbDJ4TWNnMFp4bWhmCkI2SEpWaFpXbUY5QzhsSmFJRVB3ZHRNaGQ1RWtjU1BsQlRia1NHWGpyTTQ5Y3VnbmlTSmxFc05RZ2xQYjcyM2IKRW9YQnVjNjhyUkRncFRBNTlESjNMRTlSdXdqei9ZRVNhWC92ZkZoTGVLZ0F6VnFDWEE4REM5MVozeG94Mi9mQgpaM0xiV21QeE1PSXY1NWpqSmtTaHNQajRjVWRQcEo2dDViYVFrL2pnSmNWWUo4TURXL3hiWWVIY2RKM3JVcHBNClBhNWpWRC9ZOFhvZnFESmdTNUpNalZnRkgzVkFoSjdqUUVGTzFTaWFUbzZDR3hlV2JIby9zb0Y0SjJFNTZoTWwKZ1dQN3p3S0NBUUE3dDRTK1cxTWlWZUdVZ3NZMmU0T1BXTndHMHE3RHE5R09Tek1WNUdtbmFUYVl2cnU4WHJIbAo1cy91WmJtQVR0Z0N6aHl2VVhqSEEyVjhvWUt2cnBPeUF2YjhJSlF0Ni9mYzVCMHEwT2hSa29idnJiVDdEdlpDCjJvWThnKzFWZVkzekV1SUt2ekJEQk5aSjdWTTNKRnBlaFF3UHRnMXAzWm9PL0c2YUR0S094N29QczB2Uk5YZFcKLzVjbHhkbHpIdUdOTDU4TE1hOUdseGUxNnhwOWI4TmY1OS93QUJDRzNxUkRoaVMwQjh2dXBrdlpvaUVMTE40NQpMbTJ5VFp5UmVKcG1BSUdoYkNtTlN0clM4dEhaeTh1azhNWkJOQzVKYXE4V1lTeDR5dThYcG5Cems0NWtZNWlCClMydTVCU3lOMi9ydzdGNHFVYjNGMWNuMGxpT1F4eWZuQW9JQkFGTlE0am1uWTZLQmtvVUZPTXE3TUtpUHEyM1kKMEhMdGduNTVTTlk1dUNuNmxaSGtJSTZVOTlUSWNqN24xanZFV2ZkL09oNS95T1RNWGd6VndFMUNFTXpNSVVFRgpqYlhYcWFHR3JnK1lWbFA4Vk5NdDgrc1hqRTZ4RHJVMFJTWEY0MVlySTljSVJZS0FRR3BYa0x4SzdCUHhnZm1DCm1yV2NUZVZrQlJpV1UzejVEQUJiSVV5emxpV09HNFIxZWxYQlUrWVVGOUlmUW1GcWRkSU1oVGQ3UGZWWjNOYnUKd2RxV3VGM1N1bXRzK3VBQ2VXTmJ5MVR4cnBScXMxVDMvdzNSM1VPaTJ3K0E3M1lSU2xiOEtwWUUybkEyamRZaApXeGRyMnNldVN6bjZ6dVVsTytJVDBJREVNakx6c3FDSTRzQ3JQak4rcm44U1dZVC8zVWNCRVRzY2ZFcz0KLS0tLS1FTkQgUlNBIFBSSVZBVEUgS0VZLS0tLS0K",
			//		Token:                   "93c48444c37c36b9f37fb8962591decfe3cf0507d0a4dc55beb089074aa97863330676baeb260e134bf7c338671cc4562ca6c87bc02109dad2672ef3b58abf55",
			//	},
			//},
			want: "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateKubeConf(tt.name, tt.args.serverUrl, tt.args.clientCertificateData, tt.args.clientKeyData, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateKubeConf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GenerateKubeConf() got = %v, want %v", got, tt.want)
			}
		})
	}
}
