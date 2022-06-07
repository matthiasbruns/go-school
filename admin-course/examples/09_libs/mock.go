// mocks need to be generated with mockgen

func TestGetSubscriptionStatus(t *testing.T) {
	ctrl := gomock.NewController(t) // HL1

	mockDS := mockSubscription.NewMockDataSource(ctrl) // HL1

	tested := CreateController(mockDS)
	subscriber := revenuecat.Subscriber{}

	mockDS.EXPECT().FetchSubscriberStatus("appUserId").Return(&subscriber, nil) // HL1

	status, err := tested.GetSubscriptionStatus("appUserId")

	assert.Nil(t, err)                   // HL1
	assert.Equal(t, subscriber, *status) // HL1
}