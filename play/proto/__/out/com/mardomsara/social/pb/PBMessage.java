// Code generated by Wire protocol buffer compiler, do not edit.
// Source file: all.proto
package com.mardomsara.social.pb;

import com.squareup.wire.FieldEncoding;
import com.squareup.wire.Message;
import com.squareup.wire.ProtoAdapter;
import com.squareup.wire.ProtoReader;
import com.squareup.wire.ProtoWriter;
import com.squareup.wire.WireField;
import com.squareup.wire.internal.Internal;
import java.io.IOException;
import java.lang.Integer;
import java.lang.Long;
import java.lang.Object;
import java.lang.Override;
import java.lang.String;
import java.lang.StringBuilder;
import okio.ByteString;

public final class PBMessage extends Message<PBMessage, PBMessage.Builder> {
  public static final ProtoAdapter<PBMessage> ADAPTER = new ProtoAdapter_PBMessage();

  private static final long serialVersionUID = 0L;

  public static final String DEFAULT_MESSAGEKEY = "";

  public static final String DEFAULT_ROOMKEY = "";

  public static final Integer DEFAULT_USERID = 0;

  public static final Integer DEFAULT_ROOMTYPEID = 0;

  public static final Integer DEFAULT_MESSAGETYPEID = 0;

  public static final String DEFAULT_TEXT = "";

  public static final String DEFAULT_EXTRAJSON = "";

  public static final Integer DEFAULT_ISBYME = 0;

  public static final Integer DEFAULT_ISSTARED = 0;

  public static final Long DEFAULT_CREATEDMS = 0L;

  public static final Long DEFAULT_CREATEDDEVICEMS = 0L;

  public static final Long DEFAULT_SORTID = 0L;

  public static final Long DEFAULT_PEERSEENTIME = 0L;

  public static final Long DEFAULT_SERVERRECEIVEDTIME = 0L;

  public static final Long DEFAULT_SERVERDELETEDTIME = 0L;

  public static final Long DEFAULT_ISEENTIME = 0L;

  public static final Integer DEFAULT_TOPUSH = 0;

  public static final String DEFAULT_MSGFILE_LOCALSRC = "";

  public static final Integer DEFAULT_MSGFILE_STATUS = 0;

  @WireField(
      tag = 1,
      adapter = "com.squareup.wire.ProtoAdapter#STRING"
  )
  public final String MessageKey;

  @WireField(
      tag = 2,
      adapter = "com.squareup.wire.ProtoAdapter#STRING"
  )
  public final String RoomKey;

  @WireField(
      tag = 3,
      adapter = "com.squareup.wire.ProtoAdapter#INT32"
  )
  public final Integer UserId;

  @WireField(
      tag = 4,
      adapter = "com.squareup.wire.ProtoAdapter#INT32"
  )
  public final Integer RoomTypeId;

  @WireField(
      tag = 5,
      adapter = "com.squareup.wire.ProtoAdapter#INT32"
  )
  public final Integer MessageTypeId;

  @WireField(
      tag = 6,
      adapter = "com.squareup.wire.ProtoAdapter#STRING"
  )
  public final String Text;

  @WireField(
      tag = 7,
      adapter = "com.squareup.wire.ProtoAdapter#STRING"
  )
  public final String ExtraJson;

  @WireField(
      tag = 8,
      adapter = "com.squareup.wire.ProtoAdapter#INT32"
  )
  public final Integer IsByMe;

  @WireField(
      tag = 9,
      adapter = "com.squareup.wire.ProtoAdapter#INT32"
  )
  public final Integer IsStared;

  @WireField(
      tag = 10,
      adapter = "com.squareup.wire.ProtoAdapter#INT64"
  )
  public final Long CreatedMs;

  @WireField(
      tag = 11,
      adapter = "com.squareup.wire.ProtoAdapter#INT64"
  )
  public final Long CreatedDeviceMs;

  @WireField(
      tag = 12,
      adapter = "com.squareup.wire.ProtoAdapter#INT64"
  )
  public final Long SortId;

  @WireField(
      tag = 13,
      adapter = "com.squareup.wire.ProtoAdapter#INT64"
  )
  public final Long PeerSeenTime;

  @WireField(
      tag = 14,
      adapter = "com.squareup.wire.ProtoAdapter#INT64"
  )
  public final Long ServerReceivedTime;

  @WireField(
      tag = 15,
      adapter = "com.squareup.wire.ProtoAdapter#INT64"
  )
  public final Long ServerDeletedTime;

  @WireField(
      tag = 16,
      adapter = "com.squareup.wire.ProtoAdapter#INT64"
  )
  public final Long ISeenTime;

  @WireField(
      tag = 17,
      adapter = "com.squareup.wire.ProtoAdapter#INT32"
  )
  public final Integer ToPush;

  @WireField(
      tag = 18,
      adapter = "com.squareup.wire.ProtoAdapter#STRING"
  )
  public final String MsgFile_LocalSrc;

  @WireField(
      tag = 19,
      adapter = "com.squareup.wire.ProtoAdapter#INT32"
  )
  public final Integer MsgFile_Status;

  public PBMessage(String MessageKey, String RoomKey, Integer UserId, Integer RoomTypeId,
      Integer MessageTypeId, String Text, String ExtraJson, Integer IsByMe, Integer IsStared,
      Long CreatedMs, Long CreatedDeviceMs, Long SortId, Long PeerSeenTime, Long ServerReceivedTime,
      Long ServerDeletedTime, Long ISeenTime, Integer ToPush, String MsgFile_LocalSrc,
      Integer MsgFile_Status) {
    this(MessageKey, RoomKey, UserId, RoomTypeId, MessageTypeId, Text, ExtraJson, IsByMe, IsStared, CreatedMs, CreatedDeviceMs, SortId, PeerSeenTime, ServerReceivedTime, ServerDeletedTime, ISeenTime, ToPush, MsgFile_LocalSrc, MsgFile_Status, ByteString.EMPTY);
  }

  public PBMessage(String MessageKey, String RoomKey, Integer UserId, Integer RoomTypeId,
      Integer MessageTypeId, String Text, String ExtraJson, Integer IsByMe, Integer IsStared,
      Long CreatedMs, Long CreatedDeviceMs, Long SortId, Long PeerSeenTime, Long ServerReceivedTime,
      Long ServerDeletedTime, Long ISeenTime, Integer ToPush, String MsgFile_LocalSrc,
      Integer MsgFile_Status, ByteString unknownFields) {
    super(ADAPTER, unknownFields);
    this.MessageKey = MessageKey;
    this.RoomKey = RoomKey;
    this.UserId = UserId;
    this.RoomTypeId = RoomTypeId;
    this.MessageTypeId = MessageTypeId;
    this.Text = Text;
    this.ExtraJson = ExtraJson;
    this.IsByMe = IsByMe;
    this.IsStared = IsStared;
    this.CreatedMs = CreatedMs;
    this.CreatedDeviceMs = CreatedDeviceMs;
    this.SortId = SortId;
    this.PeerSeenTime = PeerSeenTime;
    this.ServerReceivedTime = ServerReceivedTime;
    this.ServerDeletedTime = ServerDeletedTime;
    this.ISeenTime = ISeenTime;
    this.ToPush = ToPush;
    this.MsgFile_LocalSrc = MsgFile_LocalSrc;
    this.MsgFile_Status = MsgFile_Status;
  }

  @Override
  public Builder newBuilder() {
    Builder builder = new Builder();
    builder.MessageKey = MessageKey;
    builder.RoomKey = RoomKey;
    builder.UserId = UserId;
    builder.RoomTypeId = RoomTypeId;
    builder.MessageTypeId = MessageTypeId;
    builder.Text = Text;
    builder.ExtraJson = ExtraJson;
    builder.IsByMe = IsByMe;
    builder.IsStared = IsStared;
    builder.CreatedMs = CreatedMs;
    builder.CreatedDeviceMs = CreatedDeviceMs;
    builder.SortId = SortId;
    builder.PeerSeenTime = PeerSeenTime;
    builder.ServerReceivedTime = ServerReceivedTime;
    builder.ServerDeletedTime = ServerDeletedTime;
    builder.ISeenTime = ISeenTime;
    builder.ToPush = ToPush;
    builder.MsgFile_LocalSrc = MsgFile_LocalSrc;
    builder.MsgFile_Status = MsgFile_Status;
    builder.addUnknownFields(unknownFields());
    return builder;
  }

  @Override
  public boolean equals(Object other) {
    if (other == this) return true;
    if (!(other instanceof PBMessage)) return false;
    PBMessage o = (PBMessage) other;
    return unknownFields().equals(o.unknownFields())
        && Internal.equals(MessageKey, o.MessageKey)
        && Internal.equals(RoomKey, o.RoomKey)
        && Internal.equals(UserId, o.UserId)
        && Internal.equals(RoomTypeId, o.RoomTypeId)
        && Internal.equals(MessageTypeId, o.MessageTypeId)
        && Internal.equals(Text, o.Text)
        && Internal.equals(ExtraJson, o.ExtraJson)
        && Internal.equals(IsByMe, o.IsByMe)
        && Internal.equals(IsStared, o.IsStared)
        && Internal.equals(CreatedMs, o.CreatedMs)
        && Internal.equals(CreatedDeviceMs, o.CreatedDeviceMs)
        && Internal.equals(SortId, o.SortId)
        && Internal.equals(PeerSeenTime, o.PeerSeenTime)
        && Internal.equals(ServerReceivedTime, o.ServerReceivedTime)
        && Internal.equals(ServerDeletedTime, o.ServerDeletedTime)
        && Internal.equals(ISeenTime, o.ISeenTime)
        && Internal.equals(ToPush, o.ToPush)
        && Internal.equals(MsgFile_LocalSrc, o.MsgFile_LocalSrc)
        && Internal.equals(MsgFile_Status, o.MsgFile_Status);
  }

  @Override
  public int hashCode() {
    int result = super.hashCode;
    if (result == 0) {
      result = unknownFields().hashCode();
      result = result * 37 + (MessageKey != null ? MessageKey.hashCode() : 0);
      result = result * 37 + (RoomKey != null ? RoomKey.hashCode() : 0);
      result = result * 37 + (UserId != null ? UserId.hashCode() : 0);
      result = result * 37 + (RoomTypeId != null ? RoomTypeId.hashCode() : 0);
      result = result * 37 + (MessageTypeId != null ? MessageTypeId.hashCode() : 0);
      result = result * 37 + (Text != null ? Text.hashCode() : 0);
      result = result * 37 + (ExtraJson != null ? ExtraJson.hashCode() : 0);
      result = result * 37 + (IsByMe != null ? IsByMe.hashCode() : 0);
      result = result * 37 + (IsStared != null ? IsStared.hashCode() : 0);
      result = result * 37 + (CreatedMs != null ? CreatedMs.hashCode() : 0);
      result = result * 37 + (CreatedDeviceMs != null ? CreatedDeviceMs.hashCode() : 0);
      result = result * 37 + (SortId != null ? SortId.hashCode() : 0);
      result = result * 37 + (PeerSeenTime != null ? PeerSeenTime.hashCode() : 0);
      result = result * 37 + (ServerReceivedTime != null ? ServerReceivedTime.hashCode() : 0);
      result = result * 37 + (ServerDeletedTime != null ? ServerDeletedTime.hashCode() : 0);
      result = result * 37 + (ISeenTime != null ? ISeenTime.hashCode() : 0);
      result = result * 37 + (ToPush != null ? ToPush.hashCode() : 0);
      result = result * 37 + (MsgFile_LocalSrc != null ? MsgFile_LocalSrc.hashCode() : 0);
      result = result * 37 + (MsgFile_Status != null ? MsgFile_Status.hashCode() : 0);
      super.hashCode = result;
    }
    return result;
  }

  @Override
  public String toString() {
    StringBuilder builder = new StringBuilder();
    if (MessageKey != null) builder.append(", MessageKey=").append(MessageKey);
    if (RoomKey != null) builder.append(", RoomKey=").append(RoomKey);
    if (UserId != null) builder.append(", UserId=").append(UserId);
    if (RoomTypeId != null) builder.append(", RoomTypeId=").append(RoomTypeId);
    if (MessageTypeId != null) builder.append(", MessageTypeId=").append(MessageTypeId);
    if (Text != null) builder.append(", Text=").append(Text);
    if (ExtraJson != null) builder.append(", ExtraJson=").append(ExtraJson);
    if (IsByMe != null) builder.append(", IsByMe=").append(IsByMe);
    if (IsStared != null) builder.append(", IsStared=").append(IsStared);
    if (CreatedMs != null) builder.append(", CreatedMs=").append(CreatedMs);
    if (CreatedDeviceMs != null) builder.append(", CreatedDeviceMs=").append(CreatedDeviceMs);
    if (SortId != null) builder.append(", SortId=").append(SortId);
    if (PeerSeenTime != null) builder.append(", PeerSeenTime=").append(PeerSeenTime);
    if (ServerReceivedTime != null) builder.append(", ServerReceivedTime=").append(ServerReceivedTime);
    if (ServerDeletedTime != null) builder.append(", ServerDeletedTime=").append(ServerDeletedTime);
    if (ISeenTime != null) builder.append(", ISeenTime=").append(ISeenTime);
    if (ToPush != null) builder.append(", ToPush=").append(ToPush);
    if (MsgFile_LocalSrc != null) builder.append(", MsgFile_LocalSrc=").append(MsgFile_LocalSrc);
    if (MsgFile_Status != null) builder.append(", MsgFile_Status=").append(MsgFile_Status);
    return builder.replace(0, 2, "PBMessage{").append('}').toString();
  }

  public static final class Builder extends Message.Builder<PBMessage, Builder> {
    public String MessageKey;

    public String RoomKey;

    public Integer UserId;

    public Integer RoomTypeId;

    public Integer MessageTypeId;

    public String Text;

    public String ExtraJson;

    public Integer IsByMe;

    public Integer IsStared;

    public Long CreatedMs;

    public Long CreatedDeviceMs;

    public Long SortId;

    public Long PeerSeenTime;

    public Long ServerReceivedTime;

    public Long ServerDeletedTime;

    public Long ISeenTime;

    public Integer ToPush;

    public String MsgFile_LocalSrc;

    public Integer MsgFile_Status;

    public Builder() {
    }

    public Builder MessageKey(String MessageKey) {
      this.MessageKey = MessageKey;
      return this;
    }

    public Builder RoomKey(String RoomKey) {
      this.RoomKey = RoomKey;
      return this;
    }

    public Builder UserId(Integer UserId) {
      this.UserId = UserId;
      return this;
    }

    public Builder RoomTypeId(Integer RoomTypeId) {
      this.RoomTypeId = RoomTypeId;
      return this;
    }

    public Builder MessageTypeId(Integer MessageTypeId) {
      this.MessageTypeId = MessageTypeId;
      return this;
    }

    public Builder Text(String Text) {
      this.Text = Text;
      return this;
    }

    public Builder ExtraJson(String ExtraJson) {
      this.ExtraJson = ExtraJson;
      return this;
    }

    public Builder IsByMe(Integer IsByMe) {
      this.IsByMe = IsByMe;
      return this;
    }

    public Builder IsStared(Integer IsStared) {
      this.IsStared = IsStared;
      return this;
    }

    public Builder CreatedMs(Long CreatedMs) {
      this.CreatedMs = CreatedMs;
      return this;
    }

    public Builder CreatedDeviceMs(Long CreatedDeviceMs) {
      this.CreatedDeviceMs = CreatedDeviceMs;
      return this;
    }

    public Builder SortId(Long SortId) {
      this.SortId = SortId;
      return this;
    }

    public Builder PeerSeenTime(Long PeerSeenTime) {
      this.PeerSeenTime = PeerSeenTime;
      return this;
    }

    public Builder ServerReceivedTime(Long ServerReceivedTime) {
      this.ServerReceivedTime = ServerReceivedTime;
      return this;
    }

    public Builder ServerDeletedTime(Long ServerDeletedTime) {
      this.ServerDeletedTime = ServerDeletedTime;
      return this;
    }

    public Builder ISeenTime(Long ISeenTime) {
      this.ISeenTime = ISeenTime;
      return this;
    }

    public Builder ToPush(Integer ToPush) {
      this.ToPush = ToPush;
      return this;
    }

    public Builder MsgFile_LocalSrc(String MsgFile_LocalSrc) {
      this.MsgFile_LocalSrc = MsgFile_LocalSrc;
      return this;
    }

    public Builder MsgFile_Status(Integer MsgFile_Status) {
      this.MsgFile_Status = MsgFile_Status;
      return this;
    }

    @Override
    public PBMessage build() {
      return new PBMessage(MessageKey, RoomKey, UserId, RoomTypeId, MessageTypeId, Text, ExtraJson, IsByMe, IsStared, CreatedMs, CreatedDeviceMs, SortId, PeerSeenTime, ServerReceivedTime, ServerDeletedTime, ISeenTime, ToPush, MsgFile_LocalSrc, MsgFile_Status, super.buildUnknownFields());
    }
  }

  private static final class ProtoAdapter_PBMessage extends ProtoAdapter<PBMessage> {
    public ProtoAdapter_PBMessage() {
      super(FieldEncoding.LENGTH_DELIMITED, PBMessage.class);
    }

    @Override
    public int encodedSize(PBMessage value) {
      return ProtoAdapter.STRING.encodedSizeWithTag(1, value.MessageKey)
          + ProtoAdapter.STRING.encodedSizeWithTag(2, value.RoomKey)
          + ProtoAdapter.INT32.encodedSizeWithTag(3, value.UserId)
          + ProtoAdapter.INT32.encodedSizeWithTag(4, value.RoomTypeId)
          + ProtoAdapter.INT32.encodedSizeWithTag(5, value.MessageTypeId)
          + ProtoAdapter.STRING.encodedSizeWithTag(6, value.Text)
          + ProtoAdapter.STRING.encodedSizeWithTag(7, value.ExtraJson)
          + ProtoAdapter.INT32.encodedSizeWithTag(8, value.IsByMe)
          + ProtoAdapter.INT32.encodedSizeWithTag(9, value.IsStared)
          + ProtoAdapter.INT64.encodedSizeWithTag(10, value.CreatedMs)
          + ProtoAdapter.INT64.encodedSizeWithTag(11, value.CreatedDeviceMs)
          + ProtoAdapter.INT64.encodedSizeWithTag(12, value.SortId)
          + ProtoAdapter.INT64.encodedSizeWithTag(13, value.PeerSeenTime)
          + ProtoAdapter.INT64.encodedSizeWithTag(14, value.ServerReceivedTime)
          + ProtoAdapter.INT64.encodedSizeWithTag(15, value.ServerDeletedTime)
          + ProtoAdapter.INT64.encodedSizeWithTag(16, value.ISeenTime)
          + ProtoAdapter.INT32.encodedSizeWithTag(17, value.ToPush)
          + ProtoAdapter.STRING.encodedSizeWithTag(18, value.MsgFile_LocalSrc)
          + ProtoAdapter.INT32.encodedSizeWithTag(19, value.MsgFile_Status)
          + value.unknownFields().size();
    }

    @Override
    public void encode(ProtoWriter writer, PBMessage value) throws IOException {
      ProtoAdapter.STRING.encodeWithTag(writer, 1, value.MessageKey);
      ProtoAdapter.STRING.encodeWithTag(writer, 2, value.RoomKey);
      ProtoAdapter.INT32.encodeWithTag(writer, 3, value.UserId);
      ProtoAdapter.INT32.encodeWithTag(writer, 4, value.RoomTypeId);
      ProtoAdapter.INT32.encodeWithTag(writer, 5, value.MessageTypeId);
      ProtoAdapter.STRING.encodeWithTag(writer, 6, value.Text);
      ProtoAdapter.STRING.encodeWithTag(writer, 7, value.ExtraJson);
      ProtoAdapter.INT32.encodeWithTag(writer, 8, value.IsByMe);
      ProtoAdapter.INT32.encodeWithTag(writer, 9, value.IsStared);
      ProtoAdapter.INT64.encodeWithTag(writer, 10, value.CreatedMs);
      ProtoAdapter.INT64.encodeWithTag(writer, 11, value.CreatedDeviceMs);
      ProtoAdapter.INT64.encodeWithTag(writer, 12, value.SortId);
      ProtoAdapter.INT64.encodeWithTag(writer, 13, value.PeerSeenTime);
      ProtoAdapter.INT64.encodeWithTag(writer, 14, value.ServerReceivedTime);
      ProtoAdapter.INT64.encodeWithTag(writer, 15, value.ServerDeletedTime);
      ProtoAdapter.INT64.encodeWithTag(writer, 16, value.ISeenTime);
      ProtoAdapter.INT32.encodeWithTag(writer, 17, value.ToPush);
      ProtoAdapter.STRING.encodeWithTag(writer, 18, value.MsgFile_LocalSrc);
      ProtoAdapter.INT32.encodeWithTag(writer, 19, value.MsgFile_Status);
      writer.writeBytes(value.unknownFields());
    }

    @Override
    public PBMessage decode(ProtoReader reader) throws IOException {
      Builder builder = new Builder();
      long token = reader.beginMessage();
      for (int tag; (tag = reader.nextTag()) != -1;) {
        switch (tag) {
          case 1: builder.MessageKey(ProtoAdapter.STRING.decode(reader)); break;
          case 2: builder.RoomKey(ProtoAdapter.STRING.decode(reader)); break;
          case 3: builder.UserId(ProtoAdapter.INT32.decode(reader)); break;
          case 4: builder.RoomTypeId(ProtoAdapter.INT32.decode(reader)); break;
          case 5: builder.MessageTypeId(ProtoAdapter.INT32.decode(reader)); break;
          case 6: builder.Text(ProtoAdapter.STRING.decode(reader)); break;
          case 7: builder.ExtraJson(ProtoAdapter.STRING.decode(reader)); break;
          case 8: builder.IsByMe(ProtoAdapter.INT32.decode(reader)); break;
          case 9: builder.IsStared(ProtoAdapter.INT32.decode(reader)); break;
          case 10: builder.CreatedMs(ProtoAdapter.INT64.decode(reader)); break;
          case 11: builder.CreatedDeviceMs(ProtoAdapter.INT64.decode(reader)); break;
          case 12: builder.SortId(ProtoAdapter.INT64.decode(reader)); break;
          case 13: builder.PeerSeenTime(ProtoAdapter.INT64.decode(reader)); break;
          case 14: builder.ServerReceivedTime(ProtoAdapter.INT64.decode(reader)); break;
          case 15: builder.ServerDeletedTime(ProtoAdapter.INT64.decode(reader)); break;
          case 16: builder.ISeenTime(ProtoAdapter.INT64.decode(reader)); break;
          case 17: builder.ToPush(ProtoAdapter.INT32.decode(reader)); break;
          case 18: builder.MsgFile_LocalSrc(ProtoAdapter.STRING.decode(reader)); break;
          case 19: builder.MsgFile_Status(ProtoAdapter.INT32.decode(reader)); break;
          default: {
            FieldEncoding fieldEncoding = reader.peekFieldEncoding();
            Object value = fieldEncoding.rawProtoAdapter().decode(reader);
            builder.addUnknownField(tag, fieldEncoding, value);
          }
        }
      }
      reader.endMessage(token);
      return builder.build();
    }

    @Override
    public PBMessage redact(PBMessage value) {
      Builder builder = value.newBuilder();
      builder.clearUnknownFields();
      return builder.build();
    }
  }
}