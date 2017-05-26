// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: all.proto

package com.mardomsara.social.pb;

/**
 * Protobuf type {@code PBReqMsgAddMany}
 */
public  final class PBReqMsgAddMany extends
    com.google.protobuf.GeneratedMessageLite<
        PBReqMsgAddMany, PBReqMsgAddMany.Builder> implements
    // @@protoc_insertion_point(message_implements:PBReqMsgAddMany)
    PBReqMsgAddManyOrBuilder {
  private PBReqMsgAddMany() {
    messages_ = emptyProtobufList();
    users_ = emptyProtobufList();
  }
  public static final int MESSAGES_FIELD_NUMBER = 1;
  private com.google.protobuf.Internal.ProtobufList<com.mardomsara.social.pb.PBMessage> messages_;
  /**
   * <code>repeated .PBMessage Messages = 1;</code>
   */
  public java.util.List<com.mardomsara.social.pb.PBMessage> getMessagesList() {
    return messages_;
  }
  /**
   * <code>repeated .PBMessage Messages = 1;</code>
   */
  public java.util.List<? extends com.mardomsara.social.pb.PBMessageOrBuilder> 
      getMessagesOrBuilderList() {
    return messages_;
  }
  /**
   * <code>repeated .PBMessage Messages = 1;</code>
   */
  public int getMessagesCount() {
    return messages_.size();
  }
  /**
   * <code>repeated .PBMessage Messages = 1;</code>
   */
  public com.mardomsara.social.pb.PBMessage getMessages(int index) {
    return messages_.get(index);
  }
  /**
   * <code>repeated .PBMessage Messages = 1;</code>
   */
  public com.mardomsara.social.pb.PBMessageOrBuilder getMessagesOrBuilder(
      int index) {
    return messages_.get(index);
  }
  private void ensureMessagesIsMutable() {
    if (!messages_.isModifiable()) {
      messages_ =
          com.google.protobuf.GeneratedMessageLite.mutableCopy(messages_);
     }
  }

  /**
   * <code>repeated .PBMessage Messages = 1;</code>
   */
  private void setMessages(
      int index, com.mardomsara.social.pb.PBMessage value) {
    if (value == null) {
      throw new NullPointerException();
    }
    ensureMessagesIsMutable();
    messages_.set(index, value);
  }
  /**
   * <code>repeated .PBMessage Messages = 1;</code>
   */
  private void setMessages(
      int index, com.mardomsara.social.pb.PBMessage.Builder builderForValue) {
    ensureMessagesIsMutable();
    messages_.set(index, builderForValue.build());
  }
  /**
   * <code>repeated .PBMessage Messages = 1;</code>
   */
  private void addMessages(com.mardomsara.social.pb.PBMessage value) {
    if (value == null) {
      throw new NullPointerException();
    }
    ensureMessagesIsMutable();
    messages_.add(value);
  }
  /**
   * <code>repeated .PBMessage Messages = 1;</code>
   */
  private void addMessages(
      int index, com.mardomsara.social.pb.PBMessage value) {
    if (value == null) {
      throw new NullPointerException();
    }
    ensureMessagesIsMutable();
    messages_.add(index, value);
  }
  /**
   * <code>repeated .PBMessage Messages = 1;</code>
   */
  private void addMessages(
      com.mardomsara.social.pb.PBMessage.Builder builderForValue) {
    ensureMessagesIsMutable();
    messages_.add(builderForValue.build());
  }
  /**
   * <code>repeated .PBMessage Messages = 1;</code>
   */
  private void addMessages(
      int index, com.mardomsara.social.pb.PBMessage.Builder builderForValue) {
    ensureMessagesIsMutable();
    messages_.add(index, builderForValue.build());
  }
  /**
   * <code>repeated .PBMessage Messages = 1;</code>
   */
  private void addAllMessages(
      java.lang.Iterable<? extends com.mardomsara.social.pb.PBMessage> values) {
    ensureMessagesIsMutable();
    com.google.protobuf.AbstractMessageLite.addAll(
        values, messages_);
  }
  /**
   * <code>repeated .PBMessage Messages = 1;</code>
   */
  private void clearMessages() {
    messages_ = emptyProtobufList();
  }
  /**
   * <code>repeated .PBMessage Messages = 1;</code>
   */
  private void removeMessages(int index) {
    ensureMessagesIsMutable();
    messages_.remove(index);
  }

  public static final int USERS_FIELD_NUMBER = 2;
  private com.google.protobuf.Internal.ProtobufList<com.mardomsara.social.pb.PBUserWithMe> users_;
  /**
   * <code>repeated .PBUserWithMe Users = 2;</code>
   */
  public java.util.List<com.mardomsara.social.pb.PBUserWithMe> getUsersList() {
    return users_;
  }
  /**
   * <code>repeated .PBUserWithMe Users = 2;</code>
   */
  public java.util.List<? extends com.mardomsara.social.pb.PBUserWithMeOrBuilder> 
      getUsersOrBuilderList() {
    return users_;
  }
  /**
   * <code>repeated .PBUserWithMe Users = 2;</code>
   */
  public int getUsersCount() {
    return users_.size();
  }
  /**
   * <code>repeated .PBUserWithMe Users = 2;</code>
   */
  public com.mardomsara.social.pb.PBUserWithMe getUsers(int index) {
    return users_.get(index);
  }
  /**
   * <code>repeated .PBUserWithMe Users = 2;</code>
   */
  public com.mardomsara.social.pb.PBUserWithMeOrBuilder getUsersOrBuilder(
      int index) {
    return users_.get(index);
  }
  private void ensureUsersIsMutable() {
    if (!users_.isModifiable()) {
      users_ =
          com.google.protobuf.GeneratedMessageLite.mutableCopy(users_);
     }
  }

  /**
   * <code>repeated .PBUserWithMe Users = 2;</code>
   */
  private void setUsers(
      int index, com.mardomsara.social.pb.PBUserWithMe value) {
    if (value == null) {
      throw new NullPointerException();
    }
    ensureUsersIsMutable();
    users_.set(index, value);
  }
  /**
   * <code>repeated .PBUserWithMe Users = 2;</code>
   */
  private void setUsers(
      int index, com.mardomsara.social.pb.PBUserWithMe.Builder builderForValue) {
    ensureUsersIsMutable();
    users_.set(index, builderForValue.build());
  }
  /**
   * <code>repeated .PBUserWithMe Users = 2;</code>
   */
  private void addUsers(com.mardomsara.social.pb.PBUserWithMe value) {
    if (value == null) {
      throw new NullPointerException();
    }
    ensureUsersIsMutable();
    users_.add(value);
  }
  /**
   * <code>repeated .PBUserWithMe Users = 2;</code>
   */
  private void addUsers(
      int index, com.mardomsara.social.pb.PBUserWithMe value) {
    if (value == null) {
      throw new NullPointerException();
    }
    ensureUsersIsMutable();
    users_.add(index, value);
  }
  /**
   * <code>repeated .PBUserWithMe Users = 2;</code>
   */
  private void addUsers(
      com.mardomsara.social.pb.PBUserWithMe.Builder builderForValue) {
    ensureUsersIsMutable();
    users_.add(builderForValue.build());
  }
  /**
   * <code>repeated .PBUserWithMe Users = 2;</code>
   */
  private void addUsers(
      int index, com.mardomsara.social.pb.PBUserWithMe.Builder builderForValue) {
    ensureUsersIsMutable();
    users_.add(index, builderForValue.build());
  }
  /**
   * <code>repeated .PBUserWithMe Users = 2;</code>
   */
  private void addAllUsers(
      java.lang.Iterable<? extends com.mardomsara.social.pb.PBUserWithMe> values) {
    ensureUsersIsMutable();
    com.google.protobuf.AbstractMessageLite.addAll(
        values, users_);
  }
  /**
   * <code>repeated .PBUserWithMe Users = 2;</code>
   */
  private void clearUsers() {
    users_ = emptyProtobufList();
  }
  /**
   * <code>repeated .PBUserWithMe Users = 2;</code>
   */
  private void removeUsers(int index) {
    ensureUsersIsMutable();
    users_.remove(index);
  }

  public void writeTo(com.google.protobuf.CodedOutputStream output)
                      throws java.io.IOException {
    for (int i = 0; i < messages_.size(); i++) {
      output.writeMessage(1, messages_.get(i));
    }
    for (int i = 0; i < users_.size(); i++) {
      output.writeMessage(2, users_.get(i));
    }
  }

  public int getSerializedSize() {
    int size = memoizedSerializedSize;
    if (size != -1) return size;

    size = 0;
    for (int i = 0; i < messages_.size(); i++) {
      size += com.google.protobuf.CodedOutputStream
        .computeMessageSize(1, messages_.get(i));
    }
    for (int i = 0; i < users_.size(); i++) {
      size += com.google.protobuf.CodedOutputStream
        .computeMessageSize(2, users_.get(i));
    }
    memoizedSerializedSize = size;
    return size;
  }

  public static com.mardomsara.social.pb.PBReqMsgAddMany parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return com.google.protobuf.GeneratedMessageLite.parseFrom(
        DEFAULT_INSTANCE, data);
  }
  public static com.mardomsara.social.pb.PBReqMsgAddMany parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return com.google.protobuf.GeneratedMessageLite.parseFrom(
        DEFAULT_INSTANCE, data, extensionRegistry);
  }
  public static com.mardomsara.social.pb.PBReqMsgAddMany parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return com.google.protobuf.GeneratedMessageLite.parseFrom(
        DEFAULT_INSTANCE, data);
  }
  public static com.mardomsara.social.pb.PBReqMsgAddMany parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return com.google.protobuf.GeneratedMessageLite.parseFrom(
        DEFAULT_INSTANCE, data, extensionRegistry);
  }
  public static com.mardomsara.social.pb.PBReqMsgAddMany parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageLite.parseFrom(
        DEFAULT_INSTANCE, input);
  }
  public static com.mardomsara.social.pb.PBReqMsgAddMany parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageLite.parseFrom(
        DEFAULT_INSTANCE, input, extensionRegistry);
  }
  public static com.mardomsara.social.pb.PBReqMsgAddMany parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return parseDelimitedFrom(DEFAULT_INSTANCE, input);
  }
  public static com.mardomsara.social.pb.PBReqMsgAddMany parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return parseDelimitedFrom(DEFAULT_INSTANCE, input, extensionRegistry);
  }
  public static com.mardomsara.social.pb.PBReqMsgAddMany parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageLite.parseFrom(
        DEFAULT_INSTANCE, input);
  }
  public static com.mardomsara.social.pb.PBReqMsgAddMany parseFrom(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageLite.parseFrom(
        DEFAULT_INSTANCE, input, extensionRegistry);
  }

  public static Builder newBuilder() {
    return DEFAULT_INSTANCE.toBuilder();
  }
  public static Builder newBuilder(com.mardomsara.social.pb.PBReqMsgAddMany prototype) {
    return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
  }

  /**
   * Protobuf type {@code PBReqMsgAddMany}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageLite.Builder<
        com.mardomsara.social.pb.PBReqMsgAddMany, Builder> implements
      // @@protoc_insertion_point(builder_implements:PBReqMsgAddMany)
      com.mardomsara.social.pb.PBReqMsgAddManyOrBuilder {
    // Construct using com.mardomsara.social.pb.PBReqMsgAddMany.newBuilder()
    private Builder() {
      super(DEFAULT_INSTANCE);
    }


    /**
     * <code>repeated .PBMessage Messages = 1;</code>
     */
    public java.util.List<com.mardomsara.social.pb.PBMessage> getMessagesList() {
      return java.util.Collections.unmodifiableList(
          instance.getMessagesList());
    }
    /**
     * <code>repeated .PBMessage Messages = 1;</code>
     */
    public int getMessagesCount() {
      return instance.getMessagesCount();
    }/**
     * <code>repeated .PBMessage Messages = 1;</code>
     */
    public com.mardomsara.social.pb.PBMessage getMessages(int index) {
      return instance.getMessages(index);
    }
    /**
     * <code>repeated .PBMessage Messages = 1;</code>
     */
    public Builder setMessages(
        int index, com.mardomsara.social.pb.PBMessage value) {
      copyOnWrite();
      instance.setMessages(index, value);
      return this;
    }
    /**
     * <code>repeated .PBMessage Messages = 1;</code>
     */
    public Builder setMessages(
        int index, com.mardomsara.social.pb.PBMessage.Builder builderForValue) {
      copyOnWrite();
      instance.setMessages(index, builderForValue);
      return this;
    }
    /**
     * <code>repeated .PBMessage Messages = 1;</code>
     */
    public Builder addMessages(com.mardomsara.social.pb.PBMessage value) {
      copyOnWrite();
      instance.addMessages(value);
      return this;
    }
    /**
     * <code>repeated .PBMessage Messages = 1;</code>
     */
    public Builder addMessages(
        int index, com.mardomsara.social.pb.PBMessage value) {
      copyOnWrite();
      instance.addMessages(index, value);
      return this;
    }
    /**
     * <code>repeated .PBMessage Messages = 1;</code>
     */
    public Builder addMessages(
        com.mardomsara.social.pb.PBMessage.Builder builderForValue) {
      copyOnWrite();
      instance.addMessages(builderForValue);
      return this;
    }
    /**
     * <code>repeated .PBMessage Messages = 1;</code>
     */
    public Builder addMessages(
        int index, com.mardomsara.social.pb.PBMessage.Builder builderForValue) {
      copyOnWrite();
      instance.addMessages(index, builderForValue);
      return this;
    }
    /**
     * <code>repeated .PBMessage Messages = 1;</code>
     */
    public Builder addAllMessages(
        java.lang.Iterable<? extends com.mardomsara.social.pb.PBMessage> values) {
      copyOnWrite();
      instance.addAllMessages(values);
      return this;
    }
    /**
     * <code>repeated .PBMessage Messages = 1;</code>
     */
    public Builder clearMessages() {
      copyOnWrite();
      instance.clearMessages();
      return this;
    }
    /**
     * <code>repeated .PBMessage Messages = 1;</code>
     */
    public Builder removeMessages(int index) {
      copyOnWrite();
      instance.removeMessages(index);
      return this;
    }

    /**
     * <code>repeated .PBUserWithMe Users = 2;</code>
     */
    public java.util.List<com.mardomsara.social.pb.PBUserWithMe> getUsersList() {
      return java.util.Collections.unmodifiableList(
          instance.getUsersList());
    }
    /**
     * <code>repeated .PBUserWithMe Users = 2;</code>
     */
    public int getUsersCount() {
      return instance.getUsersCount();
    }/**
     * <code>repeated .PBUserWithMe Users = 2;</code>
     */
    public com.mardomsara.social.pb.PBUserWithMe getUsers(int index) {
      return instance.getUsers(index);
    }
    /**
     * <code>repeated .PBUserWithMe Users = 2;</code>
     */
    public Builder setUsers(
        int index, com.mardomsara.social.pb.PBUserWithMe value) {
      copyOnWrite();
      instance.setUsers(index, value);
      return this;
    }
    /**
     * <code>repeated .PBUserWithMe Users = 2;</code>
     */
    public Builder setUsers(
        int index, com.mardomsara.social.pb.PBUserWithMe.Builder builderForValue) {
      copyOnWrite();
      instance.setUsers(index, builderForValue);
      return this;
    }
    /**
     * <code>repeated .PBUserWithMe Users = 2;</code>
     */
    public Builder addUsers(com.mardomsara.social.pb.PBUserWithMe value) {
      copyOnWrite();
      instance.addUsers(value);
      return this;
    }
    /**
     * <code>repeated .PBUserWithMe Users = 2;</code>
     */
    public Builder addUsers(
        int index, com.mardomsara.social.pb.PBUserWithMe value) {
      copyOnWrite();
      instance.addUsers(index, value);
      return this;
    }
    /**
     * <code>repeated .PBUserWithMe Users = 2;</code>
     */
    public Builder addUsers(
        com.mardomsara.social.pb.PBUserWithMe.Builder builderForValue) {
      copyOnWrite();
      instance.addUsers(builderForValue);
      return this;
    }
    /**
     * <code>repeated .PBUserWithMe Users = 2;</code>
     */
    public Builder addUsers(
        int index, com.mardomsara.social.pb.PBUserWithMe.Builder builderForValue) {
      copyOnWrite();
      instance.addUsers(index, builderForValue);
      return this;
    }
    /**
     * <code>repeated .PBUserWithMe Users = 2;</code>
     */
    public Builder addAllUsers(
        java.lang.Iterable<? extends com.mardomsara.social.pb.PBUserWithMe> values) {
      copyOnWrite();
      instance.addAllUsers(values);
      return this;
    }
    /**
     * <code>repeated .PBUserWithMe Users = 2;</code>
     */
    public Builder clearUsers() {
      copyOnWrite();
      instance.clearUsers();
      return this;
    }
    /**
     * <code>repeated .PBUserWithMe Users = 2;</code>
     */
    public Builder removeUsers(int index) {
      copyOnWrite();
      instance.removeUsers(index);
      return this;
    }

    // @@protoc_insertion_point(builder_scope:PBReqMsgAddMany)
  }
  protected final Object dynamicMethod(
      com.google.protobuf.GeneratedMessageLite.MethodToInvoke method,
      Object arg0, Object arg1) {
    switch (method) {
      case NEW_MUTABLE_INSTANCE: {
        return new com.mardomsara.social.pb.PBReqMsgAddMany();
      }
      case IS_INITIALIZED: {
        return DEFAULT_INSTANCE;
      }
      case MAKE_IMMUTABLE: {
        messages_.makeImmutable();
        users_.makeImmutable();
        return null;
      }
      case NEW_BUILDER: {
        return new Builder();
      }
      case VISIT: {
        Visitor visitor = (Visitor) arg0;
        com.mardomsara.social.pb.PBReqMsgAddMany other = (com.mardomsara.social.pb.PBReqMsgAddMany) arg1;
        messages_= visitor.visitList(messages_, other.messages_);
        users_= visitor.visitList(users_, other.users_);
        if (visitor == com.google.protobuf.GeneratedMessageLite.MergeFromVisitor
            .INSTANCE) {
        }
        return this;
      }
      case MERGE_FROM_STREAM: {
        com.google.protobuf.CodedInputStream input =
            (com.google.protobuf.CodedInputStream) arg0;
        com.google.protobuf.ExtensionRegistryLite extensionRegistry =
            (com.google.protobuf.ExtensionRegistryLite) arg1;
        try {
          boolean done = false;
          while (!done) {
            int tag = input.readTag();
            switch (tag) {
              case 0:
                done = true;
                break;
              default: {
                if (!input.skipField(tag)) {
                  done = true;
                }
                break;
              }
              case 10: {
                if (!messages_.isModifiable()) {
                  messages_ =
                      com.google.protobuf.GeneratedMessageLite.mutableCopy(messages_);
                }
                messages_.add(
                    input.readMessage(com.mardomsara.social.pb.PBMessage.parser(), extensionRegistry));
                break;
              }
              case 18: {
                if (!users_.isModifiable()) {
                  users_ =
                      com.google.protobuf.GeneratedMessageLite.mutableCopy(users_);
                }
                users_.add(
                    input.readMessage(com.mardomsara.social.pb.PBUserWithMe.parser(), extensionRegistry));
                break;
              }
            }
          }
        } catch (com.google.protobuf.InvalidProtocolBufferException e) {
          throw new RuntimeException(e.setUnfinishedMessage(this));
        } catch (java.io.IOException e) {
          throw new RuntimeException(
              new com.google.protobuf.InvalidProtocolBufferException(
                  e.getMessage()).setUnfinishedMessage(this));
        } finally {
        }
      }
      case GET_DEFAULT_INSTANCE: {
        return DEFAULT_INSTANCE;
      }
      case GET_PARSER: {
        if (PARSER == null) {    synchronized (com.mardomsara.social.pb.PBReqMsgAddMany.class) {
            if (PARSER == null) {
              PARSER = new DefaultInstanceBasedParser(DEFAULT_INSTANCE);
            }
          }
        }
        return PARSER;
      }
    }
    throw new UnsupportedOperationException();
  }


  // @@protoc_insertion_point(class_scope:PBReqMsgAddMany)
  private static final com.mardomsara.social.pb.PBReqMsgAddMany DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new PBReqMsgAddMany();
    DEFAULT_INSTANCE.makeImmutable();
  }

  public static com.mardomsara.social.pb.PBReqMsgAddMany getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static volatile com.google.protobuf.Parser<PBReqMsgAddMany> PARSER;

  public static com.google.protobuf.Parser<PBReqMsgAddMany> parser() {
    return DEFAULT_INSTANCE.getParserForType();
  }
}
